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

type StackObservation struct {
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StackParameters struct {

	// +kubebuilder:validation:Optional
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// +kubebuilder:validation:Optional
	IamRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationArns []*string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`

	// +kubebuilder:validation:Optional
	OnFailure *string `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyBody *string `json:"policyBody,omitempty" tf:"policy_body,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyURL *string `json:"policyUrl,omitempty" tf:"policy_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutInMinutes *int64 `json:"timeoutInMinutes,omitempty" tf:"timeout_in_minutes,omitempty"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stack is the Schema for the Stacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSpec   `json:"spec"`
	Status            StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
