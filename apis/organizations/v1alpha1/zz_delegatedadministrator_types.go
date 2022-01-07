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

type DelegatedAdministratorObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DelegationEnabledDate *string `json:"delegationEnabledDate,omitempty" tf:"delegation_enabled_date,omitempty"`

	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JoinedMethod *string `json:"joinedMethod,omitempty" tf:"joined_method,omitempty"`

	JoinedTimestamp *string `json:"joinedTimestamp,omitempty" tf:"joined_timestamp,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DelegatedAdministratorParameters struct {

	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServicePrincipal *string `json:"servicePrincipal" tf:"service_principal,omitempty"`
}

// DelegatedAdministratorSpec defines the desired state of DelegatedAdministrator
type DelegatedAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DelegatedAdministratorParameters `json:"forProvider"`
}

// DelegatedAdministratorStatus defines the observed state of DelegatedAdministrator.
type DelegatedAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DelegatedAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DelegatedAdministrator is the Schema for the DelegatedAdministrators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DelegatedAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DelegatedAdministratorSpec   `json:"spec"`
	Status            DelegatedAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DelegatedAdministratorList contains a list of DelegatedAdministrators
type DelegatedAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DelegatedAdministrator `json:"items"`
}

// Repository type metadata.
var (
	DelegatedAdministrator_Kind             = "DelegatedAdministrator"
	DelegatedAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DelegatedAdministrator_Kind}.String()
	DelegatedAdministrator_KindAPIVersion   = DelegatedAdministrator_Kind + "." + CRDGroupVersion.String()
	DelegatedAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(DelegatedAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&DelegatedAdministrator{}, &DelegatedAdministratorList{})
}
