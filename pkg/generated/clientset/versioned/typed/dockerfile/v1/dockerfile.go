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

package v1

import (
	"context"
	v1 "test/pkg/apis/dockerfile/v1"
	scheme "test/pkg/generated/clientset/versioned/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DockerFilesGetter has a method to return a DockerFileInterface.
// A group's client should implement this interface.
type DockerFilesGetter interface {
	DockerFiles(namespace string) DockerFileInterface
}

// DockerFileInterface has methods to work with DockerFile resources.
type DockerFileInterface interface {
	Create(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.CreateOptions) (*v1.DockerFile, error)
	Update(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.UpdateOptions) (*v1.DockerFile, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, dockerFile *v1.DockerFile, opts metav1.UpdateOptions) (*v1.DockerFile, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DockerFile, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DockerFileList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DockerFile, err error)
	DockerFileExpansion
}

// dockerFiles implements DockerFileInterface
type dockerFiles struct {
	*gentype.ClientWithList[*v1.DockerFile, *v1.DockerFileList]
}

// newDockerFiles returns a DockerFiles
func newDockerFiles(c *HjjzsV1Client, namespace string) *dockerFiles {
	return &dockerFiles{
		gentype.NewClientWithList[*v1.DockerFile, *v1.DockerFileList](
			"dockerfiles",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.DockerFile { return &v1.DockerFile{} },
			func() *v1.DockerFileList { return &v1.DockerFileList{} }),
	}
}
