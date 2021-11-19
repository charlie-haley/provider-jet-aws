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

type EndpointObservation struct {
}

type EndpointParameters struct {

	// +kubebuilder:validation:Required
	KinesisStreamConfig []KinesisStreamConfigParameters `json:"kinesisStreamConfig" tf:"kinesis_stream_config,omitempty"`

	// +kubebuilder:validation:Required
	StreamType *string `json:"streamType" tf:"stream_type,omitempty"`
}

type KinesisStreamConfigObservation struct {
}

type KinesisStreamConfigParameters struct {

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Required
	StreamArn *string `json:"streamArn" tf:"stream_arn,omitempty"`
}

type RealtimeLogConfigObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type RealtimeLogConfigParameters struct {

	// +kubebuilder:validation:Required
	Endpoint []EndpointParameters `json:"endpoint" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Required
	Fields []*string `json:"fields" tf:"fields,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SamplingRate *int64 `json:"samplingRate" tf:"sampling_rate,omitempty"`
}

// RealtimeLogConfigSpec defines the desired state of RealtimeLogConfig
type RealtimeLogConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RealtimeLogConfigParameters `json:"forProvider"`
}

// RealtimeLogConfigStatus defines the observed state of RealtimeLogConfig.
type RealtimeLogConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RealtimeLogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RealtimeLogConfig is the Schema for the RealtimeLogConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RealtimeLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RealtimeLogConfigSpec   `json:"spec"`
	Status            RealtimeLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RealtimeLogConfigList contains a list of RealtimeLogConfigs
type RealtimeLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RealtimeLogConfig `json:"items"`
}

// Repository type metadata.
var (
	RealtimeLogConfig_Kind             = "RealtimeLogConfig"
	RealtimeLogConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RealtimeLogConfig_Kind}.String()
	RealtimeLogConfig_KindAPIVersion   = RealtimeLogConfig_Kind + "." + CRDGroupVersion.String()
	RealtimeLogConfig_GroupVersionKind = CRDGroupVersion.WithKind(RealtimeLogConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&RealtimeLogConfig{}, &RealtimeLogConfigList{})
}
