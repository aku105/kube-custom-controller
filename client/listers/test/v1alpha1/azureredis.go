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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/aku105/kube-custom-controller/apis/test/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureRedisLister helps list AzureRedises.
type AzureRedisLister interface {
	// List lists all AzureRedises in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AzureRedis, err error)
	// Get retrieves the AzureRedis from the index for a given name.
	Get(name string) (*v1alpha1.AzureRedis, error)
	AzureRedisListerExpansion
}

// azureRedisLister implements the AzureRedisLister interface.
type azureRedisLister struct {
	indexer cache.Indexer
}

// NewAzureRedisLister returns a new AzureRedisLister.
func NewAzureRedisLister(indexer cache.Indexer) AzureRedisLister {
	return &azureRedisLister{indexer: indexer}
}

// List lists all AzureRedises in the indexer.
func (s *azureRedisLister) List(selector labels.Selector) (ret []*v1alpha1.AzureRedis, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureRedis))
	})
	return ret, err
}

// Get retrieves the AzureRedis from the index for a given name.
func (s *azureRedisLister) Get(name string) (*v1alpha1.AzureRedis, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azureredis"), name)
	}
	return obj.(*v1alpha1.AzureRedis), nil
}
