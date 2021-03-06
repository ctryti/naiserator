// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	sqlcnrmcloudgooglecomv1beta1 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1beta1"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/nais/naiserator/pkg/client/listers/sql.cnrm.cloud.google.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SQLUserInformer provides access to a shared informer and lister for
// SQLUsers.
type SQLUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.SQLUserLister
}

type sQLUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSQLUserInformer constructs a new informer for SQLUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSQLUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSQLUserInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSQLUserInformer constructs a new informer for SQLUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSQLUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SqlV1beta1().SQLUsers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SqlV1beta1().SQLUsers(namespace).Watch(options)
			},
		},
		&sqlcnrmcloudgooglecomv1beta1.SQLUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *sQLUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSQLUserInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sQLUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sqlcnrmcloudgooglecomv1beta1.SQLUser{}, f.defaultInformer)
}

func (f *sQLUserInformer) Lister() v1beta1.SQLUserLister {
	return v1beta1.NewSQLUserLister(f.Informer().GetIndexer())
}
