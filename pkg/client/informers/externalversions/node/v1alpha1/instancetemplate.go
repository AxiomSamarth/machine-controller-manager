// This file was automatically generated by informer-gen

package v1alpha1

import (
	node_v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/node/v1alpha1"
	clientset "github.com/gardener/node-controller-manager/pkg/client/clientset"
	internalinterfaces "github.com/gardener/node-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/client/listers/node/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// InstanceTemplateInformer provides access to a shared informer and lister for
// InstanceTemplates.
type InstanceTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InstanceTemplateLister
}

type instanceTemplateInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewInstanceTemplateInformer constructs a new informer for InstanceTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInstanceTemplateInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.NodeV1alpha1().InstanceTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.NodeV1alpha1().InstanceTemplates(namespace).Watch(options)
			},
		},
		&node_v1alpha1.InstanceTemplate{},
		resyncPeriod,
		indexers,
	)
}

func defaultInstanceTemplateInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewInstanceTemplateInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *instanceTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&node_v1alpha1.InstanceTemplate{}, defaultInstanceTemplateInformer)
}

func (f *instanceTemplateInformer) Lister() v1alpha1.InstanceTemplateLister {
	return v1alpha1.NewInstanceTemplateLister(f.Informer().GetIndexer())
}
