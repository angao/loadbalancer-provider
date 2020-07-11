/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExtendedResourcesGetter has a method to return a ExtendedResourceInterface.
// A group's client should implement this interface.
type ExtendedResourcesGetter interface {
	ExtendedResources() ExtendedResourceInterface
}

// ExtendedResourceInterface has methods to work with ExtendedResource resources.
type ExtendedResourceInterface interface {
	Create(*v1beta1.ExtendedResource) (*v1beta1.ExtendedResource, error)
	Update(*v1beta1.ExtendedResource) (*v1beta1.ExtendedResource, error)
	UpdateStatus(*v1beta1.ExtendedResource) (*v1beta1.ExtendedResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ExtendedResource, error)
	List(opts v1.ListOptions) (*v1beta1.ExtendedResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ExtendedResource, err error)
	ExtendedResourceExpansion
}

// extendedResources implements ExtendedResourceInterface
type extendedResources struct {
	client rest.Interface
}

// newExtendedResources returns a ExtendedResources
func newExtendedResources(c *ResourceV1beta1Client) *extendedResources {
	return &extendedResources{
		client: c.RESTClient(),
	}
}

// Get takes name of the extendedResource, and returns the corresponding extendedResource object, and an error if there is any.
func (c *extendedResources) Get(name string, options v1.GetOptions) (result *v1beta1.ExtendedResource, err error) {
	result = &v1beta1.ExtendedResource{}
	err = c.client.Get().
		Resource("extendedresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExtendedResources that match those selectors.
func (c *extendedResources) List(opts v1.ListOptions) (result *v1beta1.ExtendedResourceList, err error) {
	result = &v1beta1.ExtendedResourceList{}
	err = c.client.Get().
		Resource("extendedresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested extendedResources.
func (c *extendedResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("extendedresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a extendedResource and creates it.  Returns the server's representation of the extendedResource, and an error, if there is any.
func (c *extendedResources) Create(extendedResource *v1beta1.ExtendedResource) (result *v1beta1.ExtendedResource, err error) {
	result = &v1beta1.ExtendedResource{}
	err = c.client.Post().
		Resource("extendedresources").
		Body(extendedResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a extendedResource and updates it. Returns the server's representation of the extendedResource, and an error, if there is any.
func (c *extendedResources) Update(extendedResource *v1beta1.ExtendedResource) (result *v1beta1.ExtendedResource, err error) {
	result = &v1beta1.ExtendedResource{}
	err = c.client.Put().
		Resource("extendedresources").
		Name(extendedResource.Name).
		Body(extendedResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *extendedResources) UpdateStatus(extendedResource *v1beta1.ExtendedResource) (result *v1beta1.ExtendedResource, err error) {
	result = &v1beta1.ExtendedResource{}
	err = c.client.Put().
		Resource("extendedresources").
		Name(extendedResource.Name).
		SubResource("status").
		Body(extendedResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the extendedResource and deletes it. Returns an error if one occurs.
func (c *extendedResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("extendedresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *extendedResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("extendedresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched extendedResource.
func (c *extendedResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ExtendedResource, err error) {
	result = &v1beta1.ExtendedResource{}
	err = c.client.Patch(pt).
		Resource("extendedresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
