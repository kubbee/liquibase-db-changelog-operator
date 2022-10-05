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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LiquiBaseReleaseSpec defines the desired state of LiquiBaseRelease
type LiquiBaseReleaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	SchemaId                 string                   `json:"schemaId,omitempty"`
	DBEngine                 string                   `json:"dbEngine,omitempty"`
	Hash                     string                   `json:"hash,omitempty"`
	LiquiBaseReleaseResource LiquiBaseReleaseResource `json:"resource,omitempty"`
}

// LiquiBaseReleaseResource defines the desired state of LiquidBaseRelease
type LiquiBaseReleaseResource struct {
	LiquiBaseReleaseConfigMap LiquiBaseReleaseConfigMap `json:"configMap,omitempty"`
}

// LiquiBaseReleaseConfigMap defines the desired state of LiquidBaseRelease
type LiquiBaseReleaseConfigMap struct {
	Name string `json:"name,omitempty"`
}

// LiquiBaseReleaseStatus defines the observed state of LiquiBaseRelease
type LiquiBaseReleaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ObservedGeneration int64       `json:"observedGeneration,omitempty"`
	Conditions         []Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LiquiBaseRelease is the Schema for the liquibasereleases API
type LiquiBaseRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LiquiBaseReleaseSpec   `json:"spec,omitempty"`
	Status LiquiBaseReleaseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LiquiBaseReleaseList contains a list of LiquiBaseRelease
type LiquiBaseReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiquiBaseRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiquiBaseRelease{}, &LiquiBaseReleaseList{})
}
