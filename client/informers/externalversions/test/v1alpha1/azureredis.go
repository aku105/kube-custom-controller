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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	test_v1alpha1 "github.com/amitkr0201/kube-custom-controller/apis/test/v1alpha1"
	versioned "github.com/amitkr0201/kube-custom-controller/client/clientset/versioned"
	internalinterfaces "github.com/amitkr0201/kube-custom-controller/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/amitkr0201/kube-custom-controller/client/listers/test/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureRedisInformer provides access to a shared informer and lister for
// AzureRedises.
type AzureRedisInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureRedisLister
}

type azureRedisInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzureRedisInformer constructs a new informer for AzureRedis type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureRedisInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureRedisInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzureRedisInformer constructs a new informer for AzureRedis type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureRedisInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestV1alpha1().AzureRedises().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestV1alpha1().AzureRedises().Watch(options)
			},
		},
		&test_v1alpha1.AzureRedis{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureRedisInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureRedisInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureRedisInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&test_v1alpha1.AzureRedis{}, f.defaultInformer)
}

func (f *azureRedisInformer) Lister() v1alpha1.AzureRedisLister {
	return v1alpha1.NewAzureRedisLister(f.Informer().GetIndexer())
}
