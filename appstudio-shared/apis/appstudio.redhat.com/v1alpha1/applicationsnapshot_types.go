/*
Copyright 2022.

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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplicationSnapshotSpec defines the desired state of ApplicationSnapshot
type ApplicationSnapshotSpec struct {

	// NOTE: The name (kind) of this API, "ApplicationSnapshot" is likely to change in the short term (Q2 2022).
	// Stay tuned for refactoring needed for your component.

	DisplayName        string `json:"displayName,omitempty"`
	DisplayDescription string `json:"displayDescription,omitempty"`

	Components []ApplicationSnapshotComponent `json:"components,omitempty"`
}

type ApplicationSnapshotComponent struct {
	Name           string `json:"name"`
	ContainerImage string `json:"containerImage"`
}

// Artifacts is a placeholder section for 'artifact links' we want to maintain to other AppStudio resources.
//
// For example: here I'm imagining we might want to keep track of container image <=> (source code repo, commit sha) links,
// Which might be useful to present to the user within the UI.
type Artifacts struct {

	// NOTE: This field (and struct) are placeholders.
	// - Until this API is stabilized, consumers of the API may store any unstructured JSON/YAML data here,
	//   but no backwards compatibility will be preserved.
	UnstableFields *apiextensionsv1.JSON `json:"unstableFields,omitempty"`
}

// ApplicationSnapshotStatus defines the observed state of ApplicationSnapshot
type ApplicationSnapshotStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ApplicationSnapshot is the Schema for the applicationsnapshots API
type ApplicationSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSnapshotSpec   `json:"spec,omitempty"`
	Status ApplicationSnapshotStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationSnapshotList contains a list of ApplicationSnapshot
type ApplicationSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationSnapshot{}, &ApplicationSnapshotList{})
}
