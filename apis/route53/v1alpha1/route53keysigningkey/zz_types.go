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
// +groupName=route53.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route53/v1alpha1"
)

type Route53KeySigningKeyObservation struct {
	DigestAlgorithmMnemonic string `json:"digestAlgorithmMnemonic" tf:"digest_algorithm_mnemonic"`

	DigestAlgorithmType int64 `json:"digestAlgorithmType" tf:"digest_algorithm_type"`

	DigestValue string `json:"digestValue" tf:"digest_value"`

	DnskeyRecord string `json:"dnskeyRecord" tf:"dnskey_record"`

	DsRecord string `json:"dsRecord" tf:"ds_record"`

	Flag int64 `json:"flag" tf:"flag"`

	KeyTag int64 `json:"keyTag" tf:"key_tag"`

	PublicKey string `json:"publicKey" tf:"public_key"`

	SigningAlgorithmMnemonic string `json:"signingAlgorithmMnemonic" tf:"signing_algorithm_mnemonic"`

	SigningAlgorithmType int64 `json:"signingAlgorithmType" tf:"signing_algorithm_type"`
}

type Route53KeySigningKeyParameters struct {
	HostedZoneId string `json:"hostedZoneId" tf:"hosted_zone_id"`

	KeyManagementServiceArn string `json:"keyManagementServiceArn" tf:"key_management_service_arn"`

	Name string `json:"name" tf:"name"`

	Status *string `json:"status,omitempty" tf:"status"`
}

// Route53KeySigningKeySpec defines the desired state of Route53KeySigningKey
type Route53KeySigningKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53KeySigningKeyParameters `json:"forProvider"`
}

// Route53KeySigningKeyStatus defines the observed state of Route53KeySigningKey.
type Route53KeySigningKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53KeySigningKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53KeySigningKey is the Schema for the Route53KeySigningKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53KeySigningKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53KeySigningKeySpec   `json:"spec"`
	Status            Route53KeySigningKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53KeySigningKeyList contains a list of Route53KeySigningKeys
type Route53KeySigningKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53KeySigningKey `json:"items"`
}

// Repository type metadata.
var (
	Route53KeySigningKeyKind             = "Route53KeySigningKey"
	Route53KeySigningKeyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53KeySigningKeyKind}.String()
	Route53KeySigningKeyKindAPIVersion   = Route53KeySigningKeyKind + "." + v1alpha1.GroupVersion.String()
	Route53KeySigningKeyGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53KeySigningKeyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route53KeySigningKey{}, &Route53KeySigningKeyList{})
}
