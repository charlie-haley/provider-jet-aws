/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationSnapshotObservation struct {
	ApplicationVersionID *int64 `json:"applicationVersionId,omitempty" tf:"application_version_id,omitempty"`

	SnapshotCreationTimestamp *string `json:"snapshotCreationTimestamp,omitempty" tf:"snapshot_creation_timestamp,omitempty"`
}

type ApplicationSnapshotParameters struct {

	// +kubebuilder:validation:Required
	ApplicationName *string `json:"applicationName" tf:"application_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SnapshotName *string `json:"snapshotName" tf:"snapshot_name,omitempty"`
}

// ApplicationSnapshotSpec defines the desired state of ApplicationSnapshot
type ApplicationSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationSnapshotParameters `json:"forProvider"`
}

// ApplicationSnapshotStatus defines the observed state of ApplicationSnapshot.
type ApplicationSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSnapshot is the Schema for the ApplicationSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ApplicationSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSnapshotSpec   `json:"spec"`
	Status            ApplicationSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSnapshotList contains a list of ApplicationSnapshots
type ApplicationSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSnapshot `json:"items"`
}

// Repository type metadata.
var (
	ApplicationSnapshot_Kind             = "ApplicationSnapshot"
	ApplicationSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationSnapshot_Kind}.String()
	ApplicationSnapshot_KindAPIVersion   = ApplicationSnapshot_Kind + "." + CRDGroupVersion.String()
	ApplicationSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationSnapshot{}, &ApplicationSnapshotList{})
}
