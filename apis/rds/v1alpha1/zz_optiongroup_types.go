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

type OptionGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type OptionGroupParameters struct {

	// +kubebuilder:validation:Required
	EngineName *string `json:"engineName" tf:"engine_name,omitempty"`

	// +kubebuilder:validation:Required
	MajorEngineVersion *string `json:"majorEngineVersion" tf:"major_engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Option []OptionParameters `json:"option,omitempty" tf:"option,omitempty"`

	// +kubebuilder:validation:Optional
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OptionObservation struct {
}

type OptionParameters struct {

	// +kubebuilder:validation:Optional
	DBSecurityGroupMemberships []*string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships,omitempty"`

	// +kubebuilder:validation:Required
	OptionName *string `json:"optionName" tf:"option_name,omitempty"`

	// +kubebuilder:validation:Optional
	OptionSettings []OptionSettingsParameters `json:"optionSettings,omitempty" tf:"option_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupMemberships []*string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OptionSettingsObservation struct {
}

type OptionSettingsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// OptionGroupSpec defines the desired state of OptionGroup
type OptionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OptionGroupParameters `json:"forProvider"`
}

// OptionGroupStatus defines the observed state of OptionGroup.
type OptionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OptionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OptionGroup is the Schema for the OptionGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type OptionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OptionGroupSpec   `json:"spec"`
	Status            OptionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OptionGroupList contains a list of OptionGroups
type OptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OptionGroup `json:"items"`
}

// Repository type metadata.
var (
	OptionGroup_Kind             = "OptionGroup"
	OptionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OptionGroup_Kind}.String()
	OptionGroup_KindAPIVersion   = OptionGroup_Kind + "." + CRDGroupVersion.String()
	OptionGroup_GroupVersionKind = CRDGroupVersion.WithKind(OptionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&OptionGroup{}, &OptionGroupList{})
}
