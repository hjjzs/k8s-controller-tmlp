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
	builderv1 "builder/pkg/apis/builder/v1"
	versioned "builder/pkg/client/generated/clientset/versioned"
	internalinterfaces "builder/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1 "builder/pkg/client/generated/listers/builder/v1"
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BuilderInformer provides access to a shared informer and lister for
// Builders.
type BuilderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BuilderLister
}

type builderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBuilderInformer constructs a new informer for Builder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBuilderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBuilderInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBuilderInformer constructs a new informer for Builder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBuilderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BuilderV1().Builders().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BuilderV1().Builders().Watch(context.TODO(), options)
			},
		},
		&builderv1.Builder{},
		resyncPeriod,
		indexers,
	)
}

func (f *builderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBuilderInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *builderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&builderv1.Builder{}, f.defaultInformer)
}

func (f *builderInformer) Lister() v1.BuilderLister {
	return v1.NewBuilderLister(f.Informer().GetIndexer())
}
