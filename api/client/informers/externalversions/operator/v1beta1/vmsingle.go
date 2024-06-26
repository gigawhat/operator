/*


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

package v1beta1

import (
	"context"
	time "time"

	internalinterfaces "github.com/VictoriaMetrics/operator/api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/VictoriaMetrics/operator/api/client/listers/operator/v1beta1"
	versioned "github.com/VictoriaMetrics/operator/api/client/versioned"
	operatorv1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VMSingleInformer provides access to a shared informer and lister for
// VMSingles.
type VMSingleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VMSingleLister
}

type vMSingleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVMSingleInformer constructs a new informer for VMSingle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVMSingleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVMSingleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVMSingleInformer constructs a new informer for VMSingle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVMSingleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1beta1().VMSingles(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1beta1().VMSingles(namespace).Watch(context.TODO(), options)
			},
		},
		&operatorv1beta1.VMSingle{},
		resyncPeriod,
		indexers,
	)
}

func (f *vMSingleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVMSingleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vMSingleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1beta1.VMSingle{}, f.defaultInformer)
}

func (f *vMSingleInformer) Lister() v1beta1.VMSingleLister {
	return v1beta1.NewVMSingleLister(f.Informer().GetIndexer())
}