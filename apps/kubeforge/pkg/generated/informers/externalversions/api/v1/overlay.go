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
	apiv1 "kubeforge/internal/k8s/api/v1"
	versioned "kubeforge/pkg/generated/clientset/versioned"
	internalinterfaces "kubeforge/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "kubeforge/pkg/generated/listers/api/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OverlayInformer provides access to a shared informer and lister for
// Overlays.
type OverlayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OverlayLister
}

type overlayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOverlayInformer constructs a new informer for Overlay type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOverlayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOverlayInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOverlayInformer constructs a new informer for Overlay type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOverlayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeforgeV1().Overlays(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeforgeV1().Overlays(namespace).Watch(context.TODO(), options)
			},
		},
		&apiv1.Overlay{},
		resyncPeriod,
		indexers,
	)
}

func (f *overlayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOverlayInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *overlayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiv1.Overlay{}, f.defaultInformer)
}

func (f *overlayInformer) Lister() v1.OverlayLister {
	return v1.NewOverlayLister(f.Informer().GetIndexer())
}