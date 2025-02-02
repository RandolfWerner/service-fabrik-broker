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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/osb/v1alpha1"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SFPlansGetter has a method to return a SFPlanInterface.
// A group's client should implement this interface.
type SFPlansGetter interface {
	SFPlans(namespace string) SFPlanInterface
}

// SFPlanInterface has methods to work with SFPlan resources.
type SFPlanInterface interface {
	Create(*v1alpha1.SFPlan) (*v1alpha1.SFPlan, error)
	Update(*v1alpha1.SFPlan) (*v1alpha1.SFPlan, error)
	UpdateStatus(*v1alpha1.SFPlan) (*v1alpha1.SFPlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SFPlan, error)
	List(opts v1.ListOptions) (*v1alpha1.SFPlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SFPlan, err error)
	SFPlanExpansion
}

// sFPlans implements SFPlanInterface
type sFPlans struct {
	client rest.Interface
	ns     string
}

// newSFPlans returns a SFPlans
func newSFPlans(c *OsbV1alpha1Client, namespace string) *sFPlans {
	return &sFPlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sFPlan, and returns the corresponding sFPlan object, and an error if there is any.
func (c *sFPlans) Get(name string, options v1.GetOptions) (result *v1alpha1.SFPlan, err error) {
	result = &v1alpha1.SFPlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SFPlans that match those selectors.
func (c *sFPlans) List(opts v1.ListOptions) (result *v1alpha1.SFPlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SFPlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sFPlans.
func (c *sFPlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sfplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sFPlan and creates it.  Returns the server's representation of the sFPlan, and an error, if there is any.
func (c *sFPlans) Create(sFPlan *v1alpha1.SFPlan) (result *v1alpha1.SFPlan, err error) {
	result = &v1alpha1.SFPlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sfplans").
		Body(sFPlan).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sFPlan and updates it. Returns the server's representation of the sFPlan, and an error, if there is any.
func (c *sFPlans) Update(sFPlan *v1alpha1.SFPlan) (result *v1alpha1.SFPlan, err error) {
	result = &v1alpha1.SFPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfplans").
		Name(sFPlan.Name).
		Body(sFPlan).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sFPlans) UpdateStatus(sFPlan *v1alpha1.SFPlan) (result *v1alpha1.SFPlan, err error) {
	result = &v1alpha1.SFPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfplans").
		Name(sFPlan.Name).
		SubResource("status").
		Body(sFPlan).
		Do().
		Into(result)
	return
}

// Delete takes name of the sFPlan and deletes it. Returns an error if one occurs.
func (c *sFPlans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfplans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sFPlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfplans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sFPlan.
func (c *sFPlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SFPlan, err error) {
	result = &v1alpha1.SFPlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sfplans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
