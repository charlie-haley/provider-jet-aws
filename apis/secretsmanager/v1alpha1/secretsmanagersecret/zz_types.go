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
// +groupName=secretsmanager.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/secretsmanager/v1alpha1"
)

type ReplicaObservation struct {
	LastAccessedDate string `json:"lastAccessedDate" tf:"last_accessed_date"`

	Status string `json:"status" tf:"status"`

	StatusMessage string `json:"statusMessage" tf:"status_message"`
}

type ReplicaParameters struct {
	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Region string `json:"region" tf:"region"`
}

type RotationRulesObservation struct {
}

type RotationRulesParameters struct {
	AutomaticallyAfterDays int64 `json:"automaticallyAfterDays" tf:"automatically_after_days"`
}

type SecretsmanagerSecretObservation struct {
	Arn string `json:"arn" tf:"arn"`

	RotationEnabled bool `json:"rotationEnabled" tf:"rotation_enabled"`
}

type SecretsmanagerSecretParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	ForceOverwriteReplicaSecret *bool `json:"forceOverwriteReplicaSecret,omitempty" tf:"force_overwrite_replica_secret"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Policy *string `json:"policy,omitempty" tf:"policy"`

	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`

	Replica []ReplicaParameters `json:"replica,omitempty" tf:"replica"`

	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn"`

	RotationRules []RotationRulesParameters `json:"rotationRules,omitempty" tf:"rotation_rules"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SecretsmanagerSecretSpec defines the desired state of SecretsmanagerSecret
type SecretsmanagerSecretSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecretsmanagerSecretParameters `json:"forProvider"`
}

// SecretsmanagerSecretStatus defines the observed state of SecretsmanagerSecret.
type SecretsmanagerSecretStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecretsmanagerSecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecret is the Schema for the SecretsmanagerSecrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecretsmanagerSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretSpec   `json:"spec"`
	Status            SecretsmanagerSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretList contains a list of SecretsmanagerSecrets
type SecretsmanagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsmanagerSecret `json:"items"`
}

// Repository type metadata.
var (
	SecretsmanagerSecretKind             = "SecretsmanagerSecret"
	SecretsmanagerSecretGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SecretsmanagerSecretKind}.String()
	SecretsmanagerSecretKindAPIVersion   = SecretsmanagerSecretKind + "." + v1alpha1.GroupVersion.String()
	SecretsmanagerSecretGroupVersionKind = v1alpha1.GroupVersion.WithKind(SecretsmanagerSecretKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SecretsmanagerSecret{}, &SecretsmanagerSecretList{})
}
