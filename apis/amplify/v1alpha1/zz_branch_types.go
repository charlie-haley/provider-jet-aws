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

type BranchObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AssociatedResources []*string `json:"associatedResources,omitempty" tf:"associated_resources,omitempty"`

	CustomDomains []*string `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	DestinationBranch *string `json:"destinationBranch,omitempty" tf:"destination_branch,omitempty"`

	SourceBranch *string `json:"sourceBranch,omitempty" tf:"source_branch,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type BranchParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn,omitempty"`

	// +kubebuilder:validation:Optional
	BasicAuthCredentialsSecretRef *v1.SecretKeySelector `json:"basicAuthCredentialsSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// +kubebuilder:validation:Optional
	Framework *string `json:"framework,omitempty" tf:"framework,omitempty"`

	// +kubebuilder:validation:Optional
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BranchSpec defines the desired state of Branch
type BranchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BranchParameters `json:"forProvider"`
}

// BranchStatus defines the observed state of Branch.
type BranchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BranchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Branch is the Schema for the Branchs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Branch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BranchSpec   `json:"spec"`
	Status            BranchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BranchList contains a list of Branchs
type BranchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Branch `json:"items"`
}

// Repository type metadata.
var (
	Branch_Kind             = "Branch"
	Branch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Branch_Kind}.String()
	Branch_KindAPIVersion   = Branch_Kind + "." + CRDGroupVersion.String()
	Branch_GroupVersionKind = CRDGroupVersion.WithKind(Branch_Kind)
)

func init() {
	SchemeBuilder.Register(&Branch{}, &BranchList{})
}
