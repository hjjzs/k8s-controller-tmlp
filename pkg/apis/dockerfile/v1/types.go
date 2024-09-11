package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DockerFileSpec defines the desired state of DockerFile
type DockerFileSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
	DockerFile string `json:"dockerfile"`
}

// DockerFileStatus defines the observed state of DockerFile.
// It should always be reconstructable from the state of the cluster and/or outside world.
type DockerFileStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DockerFile is the Schema for the dockerfiles API
// +k8s:openapi-gen=true
type DockerFile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerFileSpec   `json:"spec,omitempty"`
	Status DockerFileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DockerFileList contains a list of DockerFile
type DockerFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerFile `json:"items"`
}
