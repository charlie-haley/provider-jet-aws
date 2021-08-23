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
// +groupName=servicecatalog.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/servicecatalog/v1alpha1"
)

type ServicecatalogTagOptionResourceAssociationObservation struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	ResourceCreatedTime string `json:"resourceCreatedTime" tf:"resource_created_time"`

	ResourceDescription string `json:"resourceDescription" tf:"resource_description"`

	ResourceName string `json:"resourceName" tf:"resource_name"`
}

type ServicecatalogTagOptionResourceAssociationParameters struct {
	ResourceId string `json:"resourceId" tf:"resource_id"`

	TagOptionId string `json:"tagOptionId" tf:"tag_option_id"`
}

// ServicecatalogTagOptionResourceAssociationSpec defines the desired state of ServicecatalogTagOptionResourceAssociation
type ServicecatalogTagOptionResourceAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogTagOptionResourceAssociationParameters `json:"forProvider"`
}

// ServicecatalogTagOptionResourceAssociationStatus defines the observed state of ServicecatalogTagOptionResourceAssociation.
type ServicecatalogTagOptionResourceAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogTagOptionResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogTagOptionResourceAssociation is the Schema for the ServicecatalogTagOptionResourceAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogTagOptionResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogTagOptionResourceAssociationSpec   `json:"spec"`
	Status            ServicecatalogTagOptionResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogTagOptionResourceAssociationList contains a list of ServicecatalogTagOptionResourceAssociations
type ServicecatalogTagOptionResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogTagOptionResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogTagOptionResourceAssociationKind             = "ServicecatalogTagOptionResourceAssociation"
	ServicecatalogTagOptionResourceAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServicecatalogTagOptionResourceAssociationKind}.String()
	ServicecatalogTagOptionResourceAssociationKindAPIVersion   = ServicecatalogTagOptionResourceAssociationKind + "." + v1alpha1.GroupVersion.String()
	ServicecatalogTagOptionResourceAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServicecatalogTagOptionResourceAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ServicecatalogTagOptionResourceAssociation{}, &ServicecatalogTagOptionResourceAssociationList{})
}
