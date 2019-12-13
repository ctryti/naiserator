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

func (c *FakeStorageV1alpha2) GoogleStorageBuckets(namespace string) v1alpha2.GoogleStorageBucketInterface {
	return &FakeGoogleStorageBuckets{c, namespace}
}

func (c *FakeStorageV1alpha2) GoogleStorageBucketAccessControls(namespace string) v1alpha2.GoogleStorageBucketAccessControlInterface {
	return &FakeGoogleStorageBucketAccessControls{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStorageV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}