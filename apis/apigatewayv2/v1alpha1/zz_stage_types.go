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

type AccessLogSettingsObservation struct {
}

type AccessLogSettingsParameters struct {

	// +kubebuilder:validation:Required
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`
}

type DefaultRouteSettingsObservation struct {
}

type DefaultRouteSettingsParameters struct {

	// +kubebuilder:validation:Optional
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type RouteSettingsObservation struct {
}

type RouteSettingsParameters struct {

	// +kubebuilder:validation:Optional
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// +kubebuilder:validation:Required
	RouteKey *string `json:"routeKey" tf:"route_key,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type StageObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StageParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// +kubebuilder:validation:Optional
	AutoDeploy *bool `json:"autoDeploy,omitempty" tf:"auto_deploy,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRouteSettings []DefaultRouteSettingsParameters `json:"defaultRouteSettings,omitempty" tf:"default_route_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RouteSettings []RouteSettingsParameters `json:"routeSettings,omitempty" tf:"route_settings,omitempty"`

	// +kubebuilder:validation:Optional
	StageVariables map[string]*string `json:"stageVariables,omitempty" tf:"stage_variables,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StageSpec defines the desired state of Stage
type StageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StageParameters `json:"forProvider"`
}

// StageStatus defines the observed state of Stage.
type StageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stage is the Schema for the Stages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StageSpec   `json:"spec"`
	Status            StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stages
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stage `json:"items"`
}

// Repository type metadata.
var (
	Stage_Kind             = "Stage"
	Stage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stage_Kind}.String()
	Stage_KindAPIVersion   = Stage_Kind + "." + CRDGroupVersion.String()
	Stage_GroupVersionKind = CRDGroupVersion.WithKind(Stage_Kind)
)

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
