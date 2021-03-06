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

// FakeIAMServiceAccounts implements IAMServiceAccountInterface
type FakeIAMServiceAccounts struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iamserviceaccountsResource = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iamserviceaccounts"}

var iamserviceaccountsKind = schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMServiceAccount"}

// Get takes name of the iAMServiceAccount, and returns the corresponding iAMServiceAccount object, and an error if there is any.
func (c *FakeIAMServiceAccounts) Get(name string, options v1.GetOptions) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamserviceaccountsResource, c.ns, name), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// List takes label and field selectors, and returns the list of IAMServiceAccounts that match those selectors.
func (c *FakeIAMServiceAccounts) List(opts v1.ListOptions) (result *v1beta1.IAMServiceAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamserviceaccountsResource, iamserviceaccountsKind, c.ns, opts), &v1beta1.IAMServiceAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMServiceAccountList{ListMeta: obj.(*v1beta1.IAMServiceAccountList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMServiceAccounts.
func (c *FakeIAMServiceAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamserviceaccountsResource, c.ns, opts))

}

// Create takes the representation of a iAMServiceAccount and creates it.  Returns the server's representation of the iAMServiceAccount, and an error, if there is any.
func (c *FakeIAMServiceAccounts) Create(iAMServiceAccount *v1beta1.IAMServiceAccount) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamserviceaccountsResource, c.ns, iAMServiceAccount), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// Update takes the representation of a iAMServiceAccount and updates it. Returns the server's representation of the iAMServiceAccount, and an error, if there is any.
func (c *FakeIAMServiceAccounts) Update(iAMServiceAccount *v1beta1.IAMServiceAccount) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamserviceaccountsResource, c.ns, iAMServiceAccount), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// Delete takes name of the iAMServiceAccount and deletes it. Returns an error if one occurs.
func (c *FakeIAMServiceAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamserviceaccountsResource, c.ns, name), &v1beta1.IAMServiceAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMServiceAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamserviceaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMServiceAccountList{})
	return err
}

// Patch applies the patch and returns the patched iAMServiceAccount.
func (c *FakeIAMServiceAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamserviceaccountsResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}
