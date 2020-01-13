// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	iamcnrmcloudgooglecomv1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nais/naiserator/pkg/client/listers/iam.cnrm.cloud.google.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IAMPolicyMemberInformer provides access to a shared informer and lister for
// IAMPolicyMembers.
type IAMPolicyMemberInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IAMPolicyMemberLister
}

type iAMPolicyMemberInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIAMPolicyMemberInformer constructs a new informer for IAMPolicyMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIAMPolicyMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIAMPolicyMemberInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIAMPolicyMemberInformer constructs a new informer for IAMPolicyMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIAMPolicyMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1alpha1().IAMPolicyMembers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1alpha1().IAMPolicyMembers(namespace).Watch(options)
			},
		},
		&iamcnrmcloudgooglecomv1alpha1.IAMPolicyMember{},
		resyncPeriod,
		indexers,
	)
}

func (f *iAMPolicyMemberInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIAMPolicyMemberInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iAMPolicyMemberInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iamcnrmcloudgooglecomv1alpha1.IAMPolicyMember{}, f.defaultInformer)
}

func (f *iAMPolicyMemberInformer) Lister() v1alpha1.IAMPolicyMemberLister {
	return v1alpha1.NewIAMPolicyMemberLister(f.Informer().GetIndexer())
}
