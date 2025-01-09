/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceData struct{}

// ResourceModifierSpec defines the desired state of ResourceModifier.
type ResourceModifierSpec struct {
	// ResourceType field identifies the type of resource to be modified (e.g. pod, deployment, node, etc)
	ResourceType string `json:"resourceType"`

	// ResourceData will be used to identify the particular resource which user wishes to update.
	// If data specified in this field turned out to return more than 1 resource, it will result in error.
	ResourceData ResourceData `json:"resourceData"`

	// Annotations are set of pre-defined rules of how the resource will be modified.
	//
	// For example: if user has specified following annotations, and a Pod resource:
	// 	- removeAnyFinalizers
	//  - sleep:50
	// It will result in removing any finalizers Pod currently has, and executing a command to sleep for 50 seconds.
	//
	// All examples of annotations will be provided in README.
	Annotations []string `json:"annotations"`
}

// ResourceModifierStatus defines the observed state of ResourceModifier.
type ResourceModifierStatus struct {
	// Conditions are used to describe current state of ResourceModifier.
	// In case of errors, this field is updated, indicating that error had occurred.
	// If Reconciliation was successful - this fields will also be updated, with
	// successful condition type and appropriate message.
	Conditions map[string]string `json:"conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ResourceModifier is the Schema for the resourcemodifiers API.
type ResourceModifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceModifierSpec   `json:"spec,omitempty"`
	Status ResourceModifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceModifierList contains a list of ResourceModifier.
type ResourceModifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceModifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceModifier{}, &ResourceModifierList{})
}
