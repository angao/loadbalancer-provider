/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	scheme "github.com/caicloud/clientset/customclient/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeLocalStoragesGetter has a method to return a NodeLocalStorageInterface.
// A group's client should implement this interface.
type NodeLocalStoragesGetter interface {
	NodeLocalStorages() NodeLocalStorageInterface
}

// NodeLocalStorageInterface has methods to work with NodeLocalStorage resources.
type NodeLocalStorageInterface interface {
	Create(*v1beta1.NodeLocalStorage) (*v1beta1.NodeLocalStorage, error)
	Update(*v1beta1.NodeLocalStorage) (*v1beta1.NodeLocalStorage, error)
	UpdateStatus(*v1beta1.NodeLocalStorage) (*v1beta1.NodeLocalStorage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.NodeLocalStorage, error)
	List(opts v1.ListOptions) (*v1beta1.NodeLocalStorageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.NodeLocalStorage, err error)
	NodeLocalStorageExpansion
}

// nodeLocalStorages implements NodeLocalStorageInterface
type nodeLocalStorages struct {
	client rest.Interface
}

// newNodeLocalStorages returns a NodeLocalStorages
func newNodeLocalStorages(c *ResourceV1beta1Client) *nodeLocalStorages {
	return &nodeLocalStorages{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeLocalStorage, and returns the corresponding nodeLocalStorage object, and an error if there is any.
func (c *nodeLocalStorages) Get(name string, options v1.GetOptions) (result *v1beta1.NodeLocalStorage, err error) {
	result = &v1beta1.NodeLocalStorage{}
	err = c.client.Get().
		Resource("nodelocalstorages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeLocalStorages that match those selectors.
func (c *nodeLocalStorages) List(opts v1.ListOptions) (result *v1beta1.NodeLocalStorageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NodeLocalStorageList{}
	err = c.client.Get().
		Resource("nodelocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeLocalStorages.
func (c *nodeLocalStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodelocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nodeLocalStorage and creates it.  Returns the server's representation of the nodeLocalStorage, and an error, if there is any.
func (c *nodeLocalStorages) Create(nodeLocalStorage *v1beta1.NodeLocalStorage) (result *v1beta1.NodeLocalStorage, err error) {
	result = &v1beta1.NodeLocalStorage{}
	err = c.client.Post().
		Resource("nodelocalstorages").
		Body(nodeLocalStorage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeLocalStorage and updates it. Returns the server's representation of the nodeLocalStorage, and an error, if there is any.
func (c *nodeLocalStorages) Update(nodeLocalStorage *v1beta1.NodeLocalStorage) (result *v1beta1.NodeLocalStorage, err error) {
	result = &v1beta1.NodeLocalStorage{}
	err = c.client.Put().
		Resource("nodelocalstorages").
		Name(nodeLocalStorage.Name).
		Body(nodeLocalStorage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodeLocalStorages) UpdateStatus(nodeLocalStorage *v1beta1.NodeLocalStorage) (result *v1beta1.NodeLocalStorage, err error) {
	result = &v1beta1.NodeLocalStorage{}
	err = c.client.Put().
		Resource("nodelocalstorages").
		Name(nodeLocalStorage.Name).
		SubResource("status").
		Body(nodeLocalStorage).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeLocalStorage and deletes it. Returns an error if one occurs.
func (c *nodeLocalStorages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodelocalstorages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeLocalStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodelocalstorages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeLocalStorage.
func (c *nodeLocalStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.NodeLocalStorage, err error) {
	result = &v1beta1.NodeLocalStorage{}
	err = c.client.Patch(pt).
		Resource("nodelocalstorages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
