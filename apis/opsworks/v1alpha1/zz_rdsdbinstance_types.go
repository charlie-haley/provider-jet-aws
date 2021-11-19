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

type RdsDbInstanceObservation struct {
}

type RdsDbInstanceParameters struct {

	// +kubebuilder:validation:Required
	DBPasswordSecretRef v1.SecretKeySelector `json:"dbPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DBUser *string `json:"dbUser" tf:"db_user,omitempty"`

	// +kubebuilder:validation:Required
	RdsDBInstanceArn *string `json:"rdsDbInstanceArn" tf:"rds_db_instance_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StackID *string `json:"stackId" tf:"stack_id,omitempty"`
}

// RdsDbInstanceSpec defines the desired state of RdsDbInstance
type RdsDbInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdsDbInstanceParameters `json:"forProvider"`
}

// RdsDbInstanceStatus defines the observed state of RdsDbInstance.
type RdsDbInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdsDbInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RdsDbInstance is the Schema for the RdsDbInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RdsDbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsDbInstanceSpec   `json:"spec"`
	Status            RdsDbInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsDbInstanceList contains a list of RdsDbInstances
type RdsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsDbInstance `json:"items"`
}

// Repository type metadata.
var (
	RdsDbInstance_Kind             = "RdsDbInstance"
	RdsDbInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RdsDbInstance_Kind}.String()
	RdsDbInstance_KindAPIVersion   = RdsDbInstance_Kind + "." + CRDGroupVersion.String()
	RdsDbInstance_GroupVersionKind = CRDGroupVersion.WithKind(RdsDbInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RdsDbInstance{}, &RdsDbInstanceList{})
}
