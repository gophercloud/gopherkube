/*
Copyright 2024 The ORC Authors.

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

package v1alpha1

import (
	context "context"
	time "time"

	openstackresourcecontrollerapiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	clientset "github.com/k-orc/openstack-resource-controller/pkg/clients/clientset/clientset"
	internalinterfaces "github.com/k-orc/openstack-resource-controller/pkg/clients/informers/externalversions/internalinterfaces"
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/listers/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FlavorInformer provides access to a shared informer and lister for
// Flavors.
type FlavorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apiv1alpha1.FlavorLister
}

type flavorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFlavorInformer constructs a new informer for Flavor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlavorInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFlavorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFlavorInformer constructs a new informer for Flavor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlavorInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenstackV1alpha1().Flavors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenstackV1alpha1().Flavors(namespace).Watch(context.TODO(), options)
			},
		},
		&openstackresourcecontrollerapiv1alpha1.Flavor{},
		resyncPeriod,
		indexers,
	)
}

func (f *flavorInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFlavorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flavorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openstackresourcecontrollerapiv1alpha1.Flavor{}, f.defaultInformer)
}

func (f *flavorInformer) Lister() apiv1alpha1.FlavorLister {
	return apiv1alpha1.NewFlavorLister(f.Informer().GetIndexer())
}
