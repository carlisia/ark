/*
<<<<<<< HEAD
Copyright the Heptio Ark contributors.
=======
Copyright 2019 the Heptio Ark contributors.
>>>>>>> Refactor plugin structure

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

package v1

import (
	v1 "github.com/heptio/velero/pkg/apis/ark/v1"
	scheme "github.com/heptio/velero/pkg/generated/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResticRepositoriesGetter has a method to return a ResticRepositoryInterface.
// A group's client should implement this interface.
type ResticRepositoriesGetter interface {
	ResticRepositories(namespace string) ResticRepositoryInterface
}

// ResticRepositoryInterface has methods to work with ResticRepository resources.
type ResticRepositoryInterface interface {
	Create(*v1.ResticRepository) (*v1.ResticRepository, error)
	Update(*v1.ResticRepository) (*v1.ResticRepository, error)
	UpdateStatus(*v1.ResticRepository) (*v1.ResticRepository, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ResticRepository, error)
	List(opts meta_v1.ListOptions) (*v1.ResticRepositoryList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ResticRepository, err error)
	ResticRepositoryExpansion
}

// resticRepositories implements ResticRepositoryInterface
type resticRepositories struct {
	client rest.Interface
	ns     string
}

// newResticRepositories returns a ResticRepositories
func newResticRepositories(c *ArkV1Client, namespace string) *resticRepositories {
	return &resticRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resticRepository, and returns the corresponding resticRepository object, and an error if there is any.
func (c *resticRepositories) Get(name string, options meta_v1.GetOptions) (result *v1.ResticRepository, err error) {
	result = &v1.ResticRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resticrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResticRepositories that match those selectors.
func (c *resticRepositories) List(opts meta_v1.ListOptions) (result *v1.ResticRepositoryList, err error) {
	result = &v1.ResticRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resticrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resticRepositories.
func (c *resticRepositories) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resticrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a resticRepository and creates it.  Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *resticRepositories) Create(resticRepository *v1.ResticRepository) (result *v1.ResticRepository, err error) {
	result = &v1.ResticRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resticrepositories").
		Body(resticRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resticRepository and updates it. Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *resticRepositories) Update(resticRepository *v1.ResticRepository) (result *v1.ResticRepository, err error) {
	result = &v1.ResticRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resticrepositories").
		Name(resticRepository.Name).
		Body(resticRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resticRepositories) UpdateStatus(resticRepository *v1.ResticRepository) (result *v1.ResticRepository, err error) {
	result = &v1.ResticRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resticrepositories").
		Name(resticRepository.Name).
		SubResource("status").
		Body(resticRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the resticRepository and deletes it. Returns an error if one occurs.
func (c *resticRepositories) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resticrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resticRepositories) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resticrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resticRepository.
func (c *resticRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ResticRepository, err error) {
	result = &v1.ResticRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resticrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
