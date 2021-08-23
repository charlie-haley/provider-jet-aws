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

// +kubebuilder:object:generate=true
// +groupName=macie2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/macie2/v1alpha1"
)

type Macie2AccountObservation struct {
	CreatedAt string `json:"createdAt" tf:"created_at"`

	ServiceRole string `json:"serviceRole" tf:"service_role"`

	UpdatedAt string `json:"updatedAt" tf:"updated_at"`
}

type Macie2AccountParameters struct {
	FindingPublishingFrequency *string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency"`

	Status *string `json:"status,omitempty" tf:"status"`
}

// Macie2AccountSpec defines the desired state of Macie2Account
type Macie2AccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2AccountParameters `json:"forProvider"`
}

// Macie2AccountStatus defines the observed state of Macie2Account.
type Macie2AccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2Account is the Schema for the Macie2Accounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Macie2Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2AccountSpec   `json:"spec"`
	Status            Macie2AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2AccountList contains a list of Macie2Accounts
type Macie2AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2Account `json:"items"`
}

// Repository type metadata.
var (
	Macie2AccountKind             = "Macie2Account"
	Macie2AccountGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Macie2AccountKind}.String()
	Macie2AccountKindAPIVersion   = Macie2AccountKind + "." + v1alpha1.GroupVersion.String()
	Macie2AccountGroupVersionKind = v1alpha1.GroupVersion.WithKind(Macie2AccountKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Macie2Account{}, &Macie2AccountList{})
}
