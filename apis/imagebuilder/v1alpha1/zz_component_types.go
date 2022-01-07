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

type ComponentObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComponentParameters struct {

	// +kubebuilder:validation:Optional
	ChangeDescription *string `json:"changeDescription,omitempty" tf:"change_description,omitempty"`

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Platform *string `json:"platform" tf:"platform,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SupportedOsVersions []*string `json:"supportedOsVersions,omitempty" tf:"supported_os_versions,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComponentParameters `json:"forProvider"`
}

// ComponentStatus defines the observed state of Component.
type ComponentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComponentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Component is the Schema for the Components API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComponentSpec   `json:"spec"`
	Status            ComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentList contains a list of Components
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

// Repository type metadata.
var (
	Component_Kind             = "Component"
	Component_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Component_Kind}.String()
	Component_KindAPIVersion   = Component_Kind + "." + CRDGroupVersion.String()
	Component_GroupVersionKind = CRDGroupVersion.WithKind(Component_Kind)
)

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
