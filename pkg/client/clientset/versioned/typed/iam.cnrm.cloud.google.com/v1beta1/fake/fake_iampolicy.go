// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAMPolicies implements IAMPolicyInterface
type FakeIAMPolicies struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iampoliciesResource = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iampolicies"}

var iampoliciesKind = schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicy"}

// Get takes name of the iAMPolicy, and returns the corresponding iAMPolicy object, and an error if there is any.
func (c *FakeIAMPolicies) Get(name string, options v1.GetOptions) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iampoliciesResource, c.ns, name), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// List takes label and field selectors, and returns the list of IAMPolicies that match those selectors.
func (c *FakeIAMPolicies) List(opts v1.ListOptions) (result *v1beta1.IAMPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iampoliciesResource, iampoliciesKind, c.ns, opts), &v1beta1.IAMPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMPolicyList{ListMeta: obj.(*v1beta1.IAMPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMPolicies.
func (c *FakeIAMPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iampoliciesResource, c.ns, opts))

}

// Create takes the representation of a iAMPolicy and creates it.  Returns the server's representation of the iAMPolicy, and an error, if there is any.
func (c *FakeIAMPolicies) Create(iAMPolicy *v1beta1.IAMPolicy) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iampoliciesResource, c.ns, iAMPolicy), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// Update takes the representation of a iAMPolicy and updates it. Returns the server's representation of the iAMPolicy, and an error, if there is any.
func (c *FakeIAMPolicies) Update(iAMPolicy *v1beta1.IAMPolicy) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iampoliciesResource, c.ns, iAMPolicy), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// Delete takes name of the iAMPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIAMPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iampoliciesResource, c.ns, name), &v1beta1.IAMPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iAMPolicy.
func (c *FakeIAMPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iampoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}