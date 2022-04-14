/*
Copyright 2021.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CaminogoSpec defines the desired state of Caminogo
type CaminogoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Number of nodes to create. All the nodes will be created as validators
	// +optional
	// +kubebuilder:default:=5
	NodeCount int `json:"nodeCount,omitempty"`

	// Prefix,used for kubernetes objects during creation
	// +optional
	// +kubebuilder:default:="test-validator"
	DeploymentName string `json:"deploymentName,omitempty"`

	// If specified, nodes will be attached to existing network
	// +optional
	BootstrapperURL string `json:"bootstrapperURL,omitempty"`

	// Genesis for nodes, that will be attached to existing network
	// +optional
	Genesis string `json:"genesis,omitempty"`

	// Predefined secrets for nodes, quantity, should correlate to nodeCount
	// +optional
	ExistingSecrets []string `json:"existingSecrets,omitempty"`

	// Certificates for nodes, quantity, should correlate to nodeCount
	// +optional
	Certificates []Certificate `json:"certificates,omitempty"`

	// Docker image name. Will be used in chain deployments
	// +optional
	// +kubebuilder:default:="avaplatform/caminogo"
	Image string `json:"image,omitempty"`

	// Docker image tag. Will be used in chain deployments
	// +optional
	// +kubebuilder:default:="latest"
	Tag string `json:"tag,omitempty"`

	// Environment variables for caminogo.
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Resources (requests and limits of CPU and RAM) for the Caminogo instances
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	//Specify docker hub secret
	// +optional
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	//Specify Lables for Avalangego pods
	// +optional
	PodLabels map[string]string `json:"podLabels,omitempty"`

	//Specify Annotations for Avalangego pods
	// +optional
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`
}

type Certificate struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

// CaminogoStatus defines the observed state of Caminogo
type CaminogoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Service URL of the Bootstrapper node
	BootstrapperURL string `json:"bootstrapperURL"`

	// Node services list
	NetworkMembersURI []string `json:"networkMembersURI"`

	// genesis.json
	Genesis string `json:"genesis"`

	//String to indicate a logical error
	Error string `json:"error,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Caminogo is the Schema for the caminogoes API
type Caminogo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaminogoSpec   `json:"spec,omitempty"`
	Status CaminogoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CaminogoList contains a list of Caminogo
type CaminogoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Caminogo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Caminogo{}, &CaminogoList{})
}
