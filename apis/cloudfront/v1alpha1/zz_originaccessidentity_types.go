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

type OriginAccessIdentityObservation struct {
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	CloudfrontAccessIdentityPath *string `json:"cloudfrontAccessIdentityPath,omitempty" tf:"cloudfront_access_identity_path,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	IAMArn *string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	S3CanonicalUserID *string `json:"s3CanonicalUserId,omitempty" tf:"s3_canonical_user_id,omitempty"`
}

type OriginAccessIdentityParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// OriginAccessIdentitySpec defines the desired state of OriginAccessIdentity
type OriginAccessIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginAccessIdentityParameters `json:"forProvider"`
}

// OriginAccessIdentityStatus defines the observed state of OriginAccessIdentity.
type OriginAccessIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginAccessIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessIdentity is the Schema for the OriginAccessIdentitys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type OriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginAccessIdentitySpec   `json:"spec"`
	Status            OriginAccessIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessIdentityList contains a list of OriginAccessIdentitys
type OriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginAccessIdentity `json:"items"`
}

// Repository type metadata.
var (
	OriginAccessIdentity_Kind             = "OriginAccessIdentity"
	OriginAccessIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginAccessIdentity_Kind}.String()
	OriginAccessIdentity_KindAPIVersion   = OriginAccessIdentity_Kind + "." + CRDGroupVersion.String()
	OriginAccessIdentity_GroupVersionKind = CRDGroupVersion.WithKind(OriginAccessIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginAccessIdentity{}, &OriginAccessIdentityList{})
}
