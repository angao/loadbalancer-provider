/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha2 "github.com/caicloud/clientset/listers/clever/v1alpha2"
	cleverv1alpha2 "github.com/caicloud/clientset/pkg/apis/clever/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// MLNeuronInformer provides access to a shared informer and lister for
// MLNeurons.
type MLNeuronInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.MLNeuronLister
}

type mLNeuronInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMLNeuronInformer constructs a new informer for MLNeuron type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMLNeuronInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMLNeuronInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMLNeuronInformer constructs a new informer for MLNeuron type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMLNeuronInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CleverV1alpha2().MLNeurons(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CleverV1alpha2().MLNeurons(namespace).Watch(options)
			},
		},
		&cleverv1alpha2.MLNeuron{},
		resyncPeriod,
		indexers,
	)
}

func (f *mLNeuronInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMLNeuronInformer(client.(kubernetes.Interface), f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mLNeuronInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cleverv1alpha2.MLNeuron{}, f.defaultInformer)
}

func (f *mLNeuronInformer) Lister() v1alpha2.MLNeuronLister {
	return v1alpha2.NewMLNeuronLister(f.Informer().GetIndexer())
}
