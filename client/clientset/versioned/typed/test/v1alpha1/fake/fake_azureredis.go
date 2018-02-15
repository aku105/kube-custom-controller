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

package fake

import (
	v1alpha1 "github.com/aku105/kube-custom-controller/apis/test/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureRedises implements AzureRedisInterface
type FakeAzureRedises struct {
	Fake *FakeTestV1alpha1
}

var azureredisesResource = schema.GroupVersionResource{Group: "test", Version: "v1alpha1", Resource: "azureredises"}

var azureredisesKind = schema.GroupVersionKind{Group: "test", Version: "v1alpha1", Kind: "AzureRedis"}

// Get takes name of the azureRedis, and returns the corresponding azureRedis object, and an error if there is any.
func (c *FakeAzureRedises) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureRedis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azureredisesResource, name), &v1alpha1.AzureRedis{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRedis), err
}

// List takes label and field selectors, and returns the list of AzureRedises that match those selectors.
func (c *FakeAzureRedises) List(opts v1.ListOptions) (result *v1alpha1.AzureRedisList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azureredisesResource, azureredisesKind, opts), &v1alpha1.AzureRedisList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureRedisList{}
	for _, item := range obj.(*v1alpha1.AzureRedisList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureRedises.
func (c *FakeAzureRedises) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azureredisesResource, opts))
}

// Create takes the representation of a azureRedis and creates it.  Returns the server's representation of the azureRedis, and an error, if there is any.
func (c *FakeAzureRedises) Create(azureRedis *v1alpha1.AzureRedis) (result *v1alpha1.AzureRedis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azureredisesResource, azureRedis), &v1alpha1.AzureRedis{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRedis), err
}

// Update takes the representation of a azureRedis and updates it. Returns the server's representation of the azureRedis, and an error, if there is any.
func (c *FakeAzureRedises) Update(azureRedis *v1alpha1.AzureRedis) (result *v1alpha1.AzureRedis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azureredisesResource, azureRedis), &v1alpha1.AzureRedis{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRedis), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureRedises) UpdateStatus(azureRedis *v1alpha1.AzureRedis) (*v1alpha1.AzureRedis, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azureredisesResource, "status", azureRedis), &v1alpha1.AzureRedis{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRedis), err
}

// Delete takes name of the azureRedis and deletes it. Returns an error if one occurs.
func (c *FakeAzureRedises) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azureredisesResource, name), &v1alpha1.AzureRedis{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureRedises) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azureredisesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureRedisList{})
	return err
}

// Patch applies the patch and returns the patched azureRedis.
func (c *FakeAzureRedises) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRedis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azureredisesResource, name, data, subresources...), &v1alpha1.AzureRedis{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRedis), err
}
