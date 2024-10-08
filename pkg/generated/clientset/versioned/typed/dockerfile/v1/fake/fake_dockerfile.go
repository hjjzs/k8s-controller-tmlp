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
	"context"
	v1 "test/pkg/apis/dockerfile/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDockerFiles implements DockerFileInterface
type FakeDockerFiles struct {
	Fake *FakeHjjzsV1
	ns   string
}

var dockerfilesResource = v1.SchemeGroupVersion.WithResource("dockerfiles")

var dockerfilesKind = v1.SchemeGroupVersion.WithKind("DockerFile")

// Get takes name of the dockerFile, and returns the corresponding dockerFile object, and an error if there is any.
func (c *FakeDockerFiles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DockerFile, err error) {
	emptyResult := &v1.DockerFile{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(dockerfilesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DockerFile), err
}

// List takes label and field selectors, and returns the list of DockerFiles that match those selectors.
func (c *FakeDockerFiles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DockerFileList, err error) {
	emptyResult := &v1.DockerFileList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(dockerfilesResource, dockerfilesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DockerFileList{ListMeta: obj.(*v1.DockerFileList).ListMeta}
	for _, item := range obj.(*v1.DockerFileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dockerFiles.
func (c *FakeDockerFiles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(dockerfilesResource, c.ns, opts))

}

// Create takes the representation of a dockerFile and creates it.  Returns the server's representation of the dockerFile, and an error, if there is any.
func (c *FakeDockerFiles) Create(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.CreateOptions) (result *v1.DockerFile, err error) {
	emptyResult := &v1.DockerFile{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(dockerfilesResource, c.ns, dockerFile, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DockerFile), err
}

// Update takes the representation of a dockerFile and updates it. Returns the server's representation of the dockerFile, and an error, if there is any.
func (c *FakeDockerFiles) Update(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.UpdateOptions) (result *v1.DockerFile, err error) {
	emptyResult := &v1.DockerFile{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(dockerfilesResource, c.ns, dockerFile, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DockerFile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockerFiles) UpdateStatus(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.UpdateOptions) (result *v1.DockerFile, err error) {
	emptyResult := &v1.DockerFile{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(dockerfilesResource, "status", c.ns, dockerFile, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DockerFile), err
}

// Delete takes name of the dockerFile and deletes it. Returns an error if one occurs.
func (c *FakeDockerFiles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dockerfilesResource, c.ns, name, opts), &v1.DockerFile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockerFiles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(dockerfilesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DockerFileList{})
	return err
}

// Patch applies the patch and returns the patched dockerFile.
func (c *FakeDockerFiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DockerFile, err error) {
	emptyResult := &v1.DockerFile{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(dockerfilesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DockerFile), err
}
