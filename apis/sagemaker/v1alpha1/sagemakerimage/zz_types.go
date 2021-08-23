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
// +groupName=sagemaker.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/sagemaker/v1alpha1"
)

type SagemakerImageObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerImageParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	ImageName string `json:"imageName" tf:"image_name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SagemakerImageSpec defines the desired state of SagemakerImage
type SagemakerImageSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerImageParameters `json:"forProvider"`
}

// SagemakerImageStatus defines the observed state of SagemakerImage.
type SagemakerImageStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerImage is the Schema for the SagemakerImages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SagemakerImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerImageSpec   `json:"spec"`
	Status            SagemakerImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerImageList contains a list of SagemakerImages
type SagemakerImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerImage `json:"items"`
}

// Repository type metadata.
var (
	SagemakerImageKind             = "SagemakerImage"
	SagemakerImageGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SagemakerImageKind}.String()
	SagemakerImageKindAPIVersion   = SagemakerImageKind + "." + v1alpha1.GroupVersion.String()
	SagemakerImageGroupVersionKind = v1alpha1.GroupVersion.WithKind(SagemakerImageKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SagemakerImage{}, &SagemakerImageList{})
}
