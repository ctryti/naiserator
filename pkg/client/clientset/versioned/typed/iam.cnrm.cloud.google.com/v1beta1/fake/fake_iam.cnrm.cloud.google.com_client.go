// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/iam.cnrm.cloud.google.com/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIamV1beta1 struct {
	*testing.Fake
}

func (c *FakeIamV1beta1) IAMPolicies(namespace string) v1beta1.IAMPolicyInterface {
	return &FakeIAMPolicies{c, namespace}
}

func (c *FakeIamV1beta1) IAMPolicyMembers(namespace string) v1beta1.IAMPolicyMemberInterface {
	return &FakeIAMPolicyMembers{c, namespace}
}

func (c *FakeIamV1beta1) IAMServiceAccounts(namespace string) v1beta1.IAMServiceAccountInterface {
	return &FakeIAMServiceAccounts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIamV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
