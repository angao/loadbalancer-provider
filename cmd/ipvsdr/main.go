/*
Copyright 2017 Caicloud authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	corenode "github.com/caicloud/loadbalancer-provider/core/pkg/node"
	core "github.com/caicloud/loadbalancer-provider/core/provider"
	coreutil "github.com/caicloud/loadbalancer-provider/core/util"
	"github.com/caicloud/loadbalancer-provider/pkg/version"
	"github.com/caicloud/loadbalancer-provider/providers/ipvsdr"
	"gopkg.in/urfave/cli.v1"
	log "k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

// Run ...
func Run(opts *Options) error {
	info := version.Get()
	log.Infof("Provider Build Information %v", info.Pretty())

	log.Info("Provider Running with",
		"debug:", opts.Debug,
		"kubconfig:", opts.Kubeconfig,
		"lb.ns:", opts.LoadBalancerNamespace,
		"lb.name:", opts.LoadBalancerName,
		"pod.name:", opts.PodName,
		"pod.ns:", opts.PodNamespace,
	)

	log.Infof("load kubeconfig from %s", opts.Kubeconfig)
	clientset, err := coreutil.NewClientSet(opts.Kubeconfig)
	if err != nil {
		log.Fatal("Create clientset error", err)
		return err
	}

	lb, err := clientset.Custom().LoadbalanceV1alpha2().LoadBalancers(opts.LoadBalancerNamespace).Get(opts.LoadBalancerName, metav1.GetOptions{})
	if err != nil {
		log.Fatal("Can not find loadbalancer resource", "lb.ns:", opts.LoadBalancerNamespace, "lb.name:", opts.LoadBalancerName)
		return err
	}

	if lb.Spec.Providers.Ipvsdr == nil {
		return fmt.Errorf("no ipvsdr spec specified")
	}

	err = loadIPVSModule()
	if err != nil {
		log.Error("load ipvs module error", err)
		return err
	}

	nodeName, err := corenode.GetNodeNameForPod(clientset.Native(), opts.PodNamespace, opts.PodName)
	if err != nil {
		log.Fatal("Can not get node name", err)
		return err
	}

	ipvsdr, err := ipvsdr.NewIpvsdrProvider(nodeName)
	if err != nil {
		log.Errorf("Create ipvsdr provider error: %v", err)
		return err
	}

	lp := core.NewLoadBalancerProvider(&core.Configuration{
		KubeClient:            clientset,
		Backend:               ipvsdr,
		LoadBalancerName:      opts.LoadBalancerName,
		LoadBalancerNamespace: opts.LoadBalancerNamespace,
		TCPConfigMap:          lb.Status.ProxyStatus.TCPConfigMap,
		UDPConfigMap:          lb.Status.ProxyStatus.UDPConfigMap,
	})

	// handle shutdown
	go handleSigterm(lp)

	lp.Start()

	// never stop until sigterm processed
	<-wait.NeverStop

	return nil
}

func main() {
	// fix for avoiding glog Noisy logs
	_ = flag.CommandLine.Parse([]string{})

	app := cli.NewApp()
	app.Name = "provider-ipvsdr"
	app.Compiled = time.Now()
	app.Version = version.Get().Version

	// add flags to app
	opts := NewOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {
		if err := Run(opts); err != nil {
			msg := fmt.Sprintf("running loadbalancer controller failed, with err: %v\n", err)
			return cli.NewExitError(msg, 1)
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	_ = app.Run(os.Args)
}

func handleSigterm(p *core.GenericProvider) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	<-signalChan
	log.Infof("Received SIGTERM, shutting down")

	exitCode := 0
	if err := p.Stop(); err != nil {
		log.Infof("Error during shutdown %v", err)
		exitCode = 1
	}

	log.Infof("Exiting with %v", exitCode)
	os.Exit(exitCode)
}
