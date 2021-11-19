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

type UserProfileObservation struct {
}

type UserProfileParameters struct {

	// +kubebuilder:validation:Optional
	AllowSelfManagement *bool `json:"allowSelfManagement,omitempty" tf:"allow_self_management,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// +kubebuilder:validation:Required
	SSHUsername *string `json:"sshUsername" tf:"ssh_username,omitempty"`

	// +kubebuilder:validation:Required
	UserArn *string `json:"userArn" tf:"user_arn,omitempty"`
}

// UserProfileSpec defines the desired state of UserProfile
type UserProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserProfileParameters `json:"forProvider"`
}

// UserProfileStatus defines the observed state of UserProfile.
type UserProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfile is the Schema for the UserProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserProfileSpec   `json:"spec"`
	Status            UserProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfileList contains a list of UserProfiles
type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

// Repository type metadata.
var (
	UserProfile_Kind             = "UserProfile"
	UserProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserProfile_Kind}.String()
	UserProfile_KindAPIVersion   = UserProfile_Kind + "." + CRDGroupVersion.String()
	UserProfile_GroupVersionKind = CRDGroupVersion.WithKind(UserProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&UserProfile{}, &UserProfileList{})
}
