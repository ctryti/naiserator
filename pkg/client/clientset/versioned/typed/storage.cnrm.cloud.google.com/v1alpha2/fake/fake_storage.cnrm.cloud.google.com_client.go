// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/storage.cnrm.cloud.google.com/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeStorageV1alpha2 struct {
	*testing.Fake
}

func (c *FakeStorageV1alpha2) StorageBuckets(namespace string) v1alpha2.StorageBucketInterface {
	return &FakeStorageBuckets{c, namespace}
}

func (c *FakeStorageV1alpha2) StorageBucketAccessControls(namespace string) v1alpha2.StorageBucketAccessControlInterface {
	return &FakeStorageBucketAccessControls{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStorageV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
