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
// +groupName=athena.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/athena/v1alpha1"
)

type AthenaWorkgroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type AthenaWorkgroupParameters struct {
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration"`

	Description *string `json:"description,omitempty" tf:"description"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Name string `json:"name" tf:"name"`

	State *string `json:"state,omitempty" tf:"state"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {
	BytesScannedCutoffPerQuery *int64 `json:"bytesScannedCutoffPerQuery,omitempty" tf:"bytes_scanned_cutoff_per_query"`

	EnforceWorkgroupConfiguration *bool `json:"enforceWorkgroupConfiguration,omitempty" tf:"enforce_workgroup_configuration"`

	PublishCloudwatchMetricsEnabled *bool `json:"publishCloudwatchMetricsEnabled,omitempty" tf:"publish_cloudwatch_metrics_enabled"`

	RequesterPaysEnabled *bool `json:"requesterPaysEnabled,omitempty" tf:"requester_pays_enabled"`

	ResultConfiguration []ResultConfigurationParameters `json:"resultConfiguration,omitempty" tf:"result_configuration"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option"`

	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
}

type ResultConfigurationObservation struct {
}

type ResultConfigurationParameters struct {
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`

	OutputLocation *string `json:"outputLocation,omitempty" tf:"output_location"`
}

// AthenaWorkgroupSpec defines the desired state of AthenaWorkgroup
type AthenaWorkgroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AthenaWorkgroupParameters `json:"forProvider"`
}

// AthenaWorkgroupStatus defines the observed state of AthenaWorkgroup.
type AthenaWorkgroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AthenaWorkgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaWorkgroup is the Schema for the AthenaWorkgroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AthenaWorkgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaWorkgroupSpec   `json:"spec"`
	Status            AthenaWorkgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaWorkgroupList contains a list of AthenaWorkgroups
type AthenaWorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AthenaWorkgroup `json:"items"`
}

// Repository type metadata.
var (
	AthenaWorkgroupKind             = "AthenaWorkgroup"
	AthenaWorkgroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AthenaWorkgroupKind}.String()
	AthenaWorkgroupKindAPIVersion   = AthenaWorkgroupKind + "." + v1alpha1.GroupVersion.String()
	AthenaWorkgroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(AthenaWorkgroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&AthenaWorkgroup{}, &AthenaWorkgroupList{})
}
