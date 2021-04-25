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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProjectTemplateSpec defines the desired state of ProjectTemplate
type ProjectTemplateSpec struct {
	// Resources to create. Will be applied using `kubectl create -f`
	Resources string `json:"resources"`

	// TODO: allow adding metadata
}

//+kubebuilder:object:root=true

// ProjectTemplate is the Schema for the projecttemplates API
type ProjectTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectTemplateSpec   `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ProjectTemplateList contains a list of ProjectTemplate
type ProjectTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectTemplate{}, &ProjectTemplateList{})
}
