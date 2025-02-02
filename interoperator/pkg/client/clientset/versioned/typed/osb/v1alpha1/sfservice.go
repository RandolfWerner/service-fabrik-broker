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

// SFServicesGetter has a method to return a SFServiceInterface.
// A group's client should implement this interface.
type SFServicesGetter interface {
	SFServices(namespace string) SFServiceInterface
}

// SFServiceInterface has methods to work with SFService resources.
type SFServiceInterface interface {
	Create(*v1alpha1.SFService) (*v1alpha1.SFService, error)
	Update(*v1alpha1.SFService) (*v1alpha1.SFService, error)
	UpdateStatus(*v1alpha1.SFService) (*v1alpha1.SFService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SFService, error)
	List(opts v1.ListOptions) (*v1alpha1.SFServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SFService, err error)
	SFServiceExpansion
}

// sFServices implements SFServiceInterface
type sFServices struct {
	client rest.Interface
	ns     string
}

// newSFServices returns a SFServices
func newSFServices(c *OsbV1alpha1Client, namespace string) *sFServices {
	return &sFServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sFService, and returns the corresponding sFService object, and an error if there is any.
func (c *sFServices) Get(name string, options v1.GetOptions) (result *v1alpha1.SFService, err error) {
	result = &v1alpha1.SFService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SFServices that match those selectors.
func (c *sFServices) List(opts v1.ListOptions) (result *v1alpha1.SFServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SFServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sFServices.
func (c *sFServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sfservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sFService and creates it.  Returns the server's representation of the sFService, and an error, if there is any.
func (c *sFServices) Create(sFService *v1alpha1.SFService) (result *v1alpha1.SFService, err error) {
	result = &v1alpha1.SFService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sfservices").
		Body(sFService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sFService and updates it. Returns the server's representation of the sFService, and an error, if there is any.
func (c *sFServices) Update(sFService *v1alpha1.SFService) (result *v1alpha1.SFService, err error) {
	result = &v1alpha1.SFService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfservices").
		Name(sFService.Name).
		Body(sFService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sFServices) UpdateStatus(sFService *v1alpha1.SFService) (result *v1alpha1.SFService, err error) {
	result = &v1alpha1.SFService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfservices").
		Name(sFService.Name).
		SubResource("status").
		Body(sFService).
		Do().
		Into(result)
	return
}

// Delete takes name of the sFService and deletes it. Returns an error if one occurs.
func (c *sFServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sFServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sFService.
func (c *sFServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SFService, err error) {
	result = &v1alpha1.SFService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sfservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
