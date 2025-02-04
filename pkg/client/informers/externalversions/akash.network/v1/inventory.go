/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	akashnetworkv1 "github.com/ovrclk/akash/pkg/apis/akash.network/v1"
	versioned "github.com/ovrclk/akash/pkg/client/clientset/versioned"
	internalinterfaces "github.com/ovrclk/akash/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/ovrclk/akash/pkg/client/listers/akash.network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InventoryInformer provides access to a shared informer and lister for
// Inventories.
type InventoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.InventoryLister
}

type inventoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewInventoryInformer constructs a new informer for Inventory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInventoryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInventoryInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredInventoryInformer constructs a new informer for Inventory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInventoryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AkashV1().Inventories().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AkashV1().Inventories().Watch(context.TODO(), options)
			},
		},
		&akashnetworkv1.Inventory{},
		resyncPeriod,
		indexers,
	)
}

func (f *inventoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInventoryInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inventoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&akashnetworkv1.Inventory{}, f.defaultInformer)
}

func (f *inventoryInformer) Lister() v1.InventoryLister {
	return v1.NewInventoryLister(f.Informer().GetIndexer())
}
