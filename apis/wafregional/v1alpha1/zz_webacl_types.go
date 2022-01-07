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

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DefaultActionObservation struct {
}

type DefaultActionParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LoggingConfigurationObservation struct {
}

type LoggingConfigurationParameters struct {

	// +kubebuilder:validation:Required
	LogDestination *string `json:"logDestination" tf:"log_destination,omitempty"`

	// +kubebuilder:validation:Optional
	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type OverrideActionObservation struct {
}

type OverrideActionParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedactedFieldsFieldToMatchObservation struct {
}

type RedactedFieldsFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedactedFieldsObservation struct {
}

type RedactedFieldsParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []RedactedFieldsFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`
}

type WebACLObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WebACLParameters struct {

	// +kubebuilder:validation:Required
	DefaultAction []DefaultActionParameters `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Rule []WebACLRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebACLRuleObservation struct {
}

type WebACLRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	OverrideAction []OverrideActionParameters `json:"overrideAction,omitempty" tf:"override_action,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	RuleID *string `json:"ruleId" tf:"rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// WebACLSpec defines the desired state of WebACL
type WebACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebACLParameters `json:"forProvider"`
}

// WebACLStatus defines the observed state of WebACL.
type WebACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebACL is the Schema for the WebACLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type WebACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebACLSpec   `json:"spec"`
	Status            WebACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLList contains a list of WebACLs
type WebACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebACL `json:"items"`
}

// Repository type metadata.
var (
	WebACL_Kind             = "WebACL"
	WebACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebACL_Kind}.String()
	WebACL_KindAPIVersion   = WebACL_Kind + "." + CRDGroupVersion.String()
	WebACL_GroupVersionKind = CRDGroupVersion.WithKind(WebACL_Kind)
)

func init() {
	SchemeBuilder.Register(&WebACL{}, &WebACLList{})
}
