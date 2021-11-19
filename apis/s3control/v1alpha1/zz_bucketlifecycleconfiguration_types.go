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

type AbortIncompleteMultipartUploadObservation struct {
}

type AbortIncompleteMultipartUploadParameters struct {

	// +kubebuilder:validation:Required
	DaysAfterInitiation *int64 `json:"daysAfterInitiation" tf:"days_after_initiation,omitempty"`
}

type BucketLifecycleConfigurationObservation struct {
}

type BucketLifecycleConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// BucketLifecycleConfigurationSpec defines the desired state of BucketLifecycleConfiguration
type BucketLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleConfigurationParameters `json:"forProvider"`
}

// BucketLifecycleConfigurationStatus defines the observed state of BucketLifecycleConfiguration.
type BucketLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfiguration is the Schema for the BucketLifecycleConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketLifecycleConfigurationSpec   `json:"spec"`
	Status            BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfigurationList contains a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycleConfiguration_Kind             = "BucketLifecycleConfiguration"
	BucketLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycleConfiguration_Kind}.String()
	BucketLifecycleConfiguration_KindAPIVersion   = BucketLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycleConfiguration{}, &BucketLifecycleConfigurationList{})
}
