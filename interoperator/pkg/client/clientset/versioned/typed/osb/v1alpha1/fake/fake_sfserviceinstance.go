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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/osb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSFServiceInstances implements SFServiceInstanceInterface
type FakeSFServiceInstances struct {
	Fake *FakeOsbV1alpha1
	ns   string
}

var sfserviceinstancesResource = schema.GroupVersionResource{Group: "osb.servicefabrik.io", Version: "v1alpha1", Resource: "sfserviceinstances"}

var sfserviceinstancesKind = schema.GroupVersionKind{Group: "osb.servicefabrik.io", Version: "v1alpha1", Kind: "SFServiceInstance"}

// Get takes name of the sFServiceInstance, and returns the corresponding sFServiceInstance object, and an error if there is any.
func (c *FakeSFServiceInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.SFServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sfserviceinstancesResource, c.ns, name), &v1alpha1.SFServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SFServiceInstance), err
}

// List takes label and field selectors, and returns the list of SFServiceInstances that match those selectors.
func (c *FakeSFServiceInstances) List(opts v1.ListOptions) (result *v1alpha1.SFServiceInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sfserviceinstancesResource, sfserviceinstancesKind, c.ns, opts), &v1alpha1.SFServiceInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SFServiceInstanceList{ListMeta: obj.(*v1alpha1.SFServiceInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SFServiceInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sFServiceInstances.
func (c *FakeSFServiceInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sfserviceinstancesResource, c.ns, opts))

}

// Create takes the representation of a sFServiceInstance and creates it.  Returns the server's representation of the sFServiceInstance, and an error, if there is any.
func (c *FakeSFServiceInstances) Create(sFServiceInstance *v1alpha1.SFServiceInstance) (result *v1alpha1.SFServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sfserviceinstancesResource, c.ns, sFServiceInstance), &v1alpha1.SFServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SFServiceInstance), err
}

// Update takes the representation of a sFServiceInstance and updates it. Returns the server's representation of the sFServiceInstance, and an error, if there is any.
func (c *FakeSFServiceInstances) Update(sFServiceInstance *v1alpha1.SFServiceInstance) (result *v1alpha1.SFServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sfserviceinstancesResource, c.ns, sFServiceInstance), &v1alpha1.SFServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SFServiceInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSFServiceInstances) UpdateStatus(sFServiceInstance *v1alpha1.SFServiceInstance) (*v1alpha1.SFServiceInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sfserviceinstancesResource, "status", c.ns, sFServiceInstance), &v1alpha1.SFServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SFServiceInstance), err
}

// Delete takes name of the sFServiceInstance and deletes it. Returns an error if one occurs.
func (c *FakeSFServiceInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sfserviceinstancesResource, c.ns, name), &v1alpha1.SFServiceInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSFServiceInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sfserviceinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SFServiceInstanceList{})
	return err
}

// Patch applies the patch and returns the patched sFServiceInstance.
func (c *FakeSFServiceInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SFServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sfserviceinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SFServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SFServiceInstance), err
}
