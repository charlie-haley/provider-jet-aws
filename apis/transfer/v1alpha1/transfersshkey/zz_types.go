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
// +groupName=transfer.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/transfer/v1alpha1"
)

type TransferSshKeyObservation struct {
}

type TransferSshKeyParameters struct {
	Body string `json:"body" tf:"body"`

	ServerId string `json:"serverId" tf:"server_id"`

	UserName string `json:"userName" tf:"user_name"`
}

// TransferSshKeySpec defines the desired state of TransferSshKey
type TransferSshKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TransferSshKeyParameters `json:"forProvider"`
}

// TransferSshKeyStatus defines the observed state of TransferSshKey.
type TransferSshKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TransferSshKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransferSshKey is the Schema for the TransferSshKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransferSshKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferSshKeySpec   `json:"spec"`
	Status            TransferSshKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferSshKeyList contains a list of TransferSshKeys
type TransferSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferSshKey `json:"items"`
}

// Repository type metadata.
var (
	TransferSshKeyKind             = "TransferSshKey"
	TransferSshKeyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: TransferSshKeyKind}.String()
	TransferSshKeyKindAPIVersion   = TransferSshKeyKind + "." + v1alpha1.GroupVersion.String()
	TransferSshKeyGroupVersionKind = v1alpha1.GroupVersion.WithKind(TransferSshKeyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&TransferSshKey{}, &TransferSshKeyList{})
}
