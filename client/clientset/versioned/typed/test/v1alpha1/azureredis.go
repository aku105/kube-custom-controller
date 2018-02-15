/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/amitkr0201/kube-custom-controller/apis/test/v1alpha1"
	scheme "github.com/amitkr0201/kube-custom-controller/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureRedisesGetter has a method to return a AzureRedisInterface.
// A group's client should implement this interface.
type AzureRedisesGetter interface {
	AzureRedises() AzureRedisInterface
}

// AzureRedisInterface has methods to work with AzureRedis resources.
type AzureRedisInterface interface {
	Create(*v1alpha1.AzureRedis) (*v1alpha1.AzureRedis, error)
	Update(*v1alpha1.AzureRedis) (*v1alpha1.AzureRedis, error)
	UpdateStatus(*v1alpha1.AzureRedis) (*v1alpha1.AzureRedis, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzureRedis, error)
	List(opts v1.ListOptions) (*v1alpha1.AzureRedisList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRedis, err error)
	AzureRedisExpansion
}

// azureRedises implements AzureRedisInterface
type azureRedises struct {
	client rest.Interface
}

// newAzureRedises returns a AzureRedises
func newAzureRedises(c *TestV1alpha1Client) *azureRedises {
	return &azureRedises{
		client: c.RESTClient(),
	}
}

// Get takes name of the azureRedis, and returns the corresponding azureRedis object, and an error if there is any.
func (c *azureRedises) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureRedis, err error) {
	result = &v1alpha1.AzureRedis{}
	err = c.client.Get().
		Resource("azureredises").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureRedises that match those selectors.
func (c *azureRedises) List(opts v1.ListOptions) (result *v1alpha1.AzureRedisList, err error) {
	result = &v1alpha1.AzureRedisList{}
	err = c.client.Get().
		Resource("azureredises").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureRedises.
func (c *azureRedises) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("azureredises").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a azureRedis and creates it.  Returns the server's representation of the azureRedis, and an error, if there is any.
func (c *azureRedises) Create(azureRedis *v1alpha1.AzureRedis) (result *v1alpha1.AzureRedis, err error) {
	result = &v1alpha1.AzureRedis{}
	err = c.client.Post().
		Resource("azureredises").
		Body(azureRedis).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureRedis and updates it. Returns the server's representation of the azureRedis, and an error, if there is any.
func (c *azureRedises) Update(azureRedis *v1alpha1.AzureRedis) (result *v1alpha1.AzureRedis, err error) {
	result = &v1alpha1.AzureRedis{}
	err = c.client.Put().
		Resource("azureredises").
		Name(azureRedis.Name).
		Body(azureRedis).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azureRedises) UpdateStatus(azureRedis *v1alpha1.AzureRedis) (result *v1alpha1.AzureRedis, err error) {
	result = &v1alpha1.AzureRedis{}
	err = c.client.Put().
		Resource("azureredises").
		Name(azureRedis.Name).
		SubResource("status").
		Body(azureRedis).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureRedis and deletes it. Returns an error if one occurs.
func (c *azureRedises) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azureredises").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureRedises) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("azureredises").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureRedis.
func (c *azureRedises) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRedis, err error) {
	result = &v1alpha1.AzureRedis{}
	err = c.client.Patch(pt).
		Resource("azureredises").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
