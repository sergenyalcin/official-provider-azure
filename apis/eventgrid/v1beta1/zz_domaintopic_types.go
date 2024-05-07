// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainTopicInitParameters struct {
}

type DomainTopicObservation struct {

	// Specifies the name of the EventGrid Domain. Changing this forces a new resource to be created.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The ID of the EventGrid Domain Topic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type DomainTopicParameters struct {

	// Specifies the name of the EventGrid Domain. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventgrid/v1beta2.Domain
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in eventgrid to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in eventgrid to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// DomainTopicSpec defines the desired state of DomainTopic
type DomainTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainTopicParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DomainTopicInitParameters `json:"initProvider,omitempty"`
}

// DomainTopicStatus defines the observed state of DomainTopic.
type DomainTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainTopic is the Schema for the DomainTopics API. Manages an EventGrid Domain Topic
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DomainTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainTopicSpec   `json:"spec"`
	Status            DomainTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainTopicList contains a list of DomainTopics
type DomainTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainTopic `json:"items"`
}

// Repository type metadata.
var (
	DomainTopic_Kind             = "DomainTopic"
	DomainTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainTopic_Kind}.String()
	DomainTopic_KindAPIVersion   = DomainTopic_Kind + "." + CRDGroupVersion.String()
	DomainTopic_GroupVersionKind = CRDGroupVersion.WithKind(DomainTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainTopic{}, &DomainTopicList{})
}
