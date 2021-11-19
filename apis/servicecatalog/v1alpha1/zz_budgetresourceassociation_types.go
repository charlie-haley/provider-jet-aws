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

type BudgetResourceAssociationObservation struct {
}

type BudgetResourceAssociationParameters struct {

	// +kubebuilder:validation:Required
	BudgetName *string `json:"budgetName" tf:"budget_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`
}

// BudgetResourceAssociationSpec defines the desired state of BudgetResourceAssociation
type BudgetResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetResourceAssociationParameters `json:"forProvider"`
}

// BudgetResourceAssociationStatus defines the observed state of BudgetResourceAssociation.
type BudgetResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetResourceAssociation is the Schema for the BudgetResourceAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BudgetResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetResourceAssociationSpec   `json:"spec"`
	Status            BudgetResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetResourceAssociationList contains a list of BudgetResourceAssociations
type BudgetResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	BudgetResourceAssociation_Kind             = "BudgetResourceAssociation"
	BudgetResourceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetResourceAssociation_Kind}.String()
	BudgetResourceAssociation_KindAPIVersion   = BudgetResourceAssociation_Kind + "." + CRDGroupVersion.String()
	BudgetResourceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(BudgetResourceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetResourceAssociation{}, &BudgetResourceAssociationList{})
}
