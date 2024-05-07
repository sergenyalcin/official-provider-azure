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

type IntegrationRuntimeAzureInitParameters struct {

	// Compute type of the cluster which will execute data flow job. Valid values are General, ComputeOptimized and MemoryOptimized. Defaults to General.
	ComputeType *string `json:"computeType,omitempty" tf:"compute_type,omitempty"`

	// Core count of the cluster which will execute data flow job. Valid values are 8, 16, 32, 48, 80, 144 and 272. Defaults to 8.
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Azure Region where the Synapse Azure Integration Runtime should exist. Use AutoResolve to create an auto-resolve integration runtime. Changing this forces a new Synapse Azure Integration Runtime to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to 0.
	TimeToLiveMin *float64 `json:"timeToLiveMin,omitempty" tf:"time_to_live_min,omitempty"`
}

type IntegrationRuntimeAzureObservation struct {

	// Compute type of the cluster which will execute data flow job. Valid values are General, ComputeOptimized and MemoryOptimized. Defaults to General.
	ComputeType *string `json:"computeType,omitempty" tf:"compute_type,omitempty"`

	// Core count of the cluster which will execute data flow job. Valid values are 8, 16, 32, 48, 80, 144 and 272. Defaults to 8.
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Synapse Azure Integration Runtime.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Synapse Azure Integration Runtime should exist. Use AutoResolve to create an auto-resolve integration runtime. Changing this forces a new Synapse Azure Integration Runtime to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Synapse Workspace ID in which to associate the Integration Runtime with. Changing this forces a new Synapse Azure Integration Runtime to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to 0.
	TimeToLiveMin *float64 `json:"timeToLiveMin,omitempty" tf:"time_to_live_min,omitempty"`
}

type IntegrationRuntimeAzureParameters struct {

	// Compute type of the cluster which will execute data flow job. Valid values are General, ComputeOptimized and MemoryOptimized. Defaults to General.
	// +kubebuilder:validation:Optional
	ComputeType *string `json:"computeType,omitempty" tf:"compute_type,omitempty"`

	// Core count of the cluster which will execute data flow job. Valid values are 8, 16, 32, 48, 80, 144 and 272. Defaults to 8.
	// +kubebuilder:validation:Optional
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// Integration runtime description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Azure Region where the Synapse Azure Integration Runtime should exist. Use AutoResolve to create an auto-resolve integration runtime. Changing this forces a new Synapse Azure Integration Runtime to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Synapse Workspace ID in which to associate the Integration Runtime with. Changing this forces a new Synapse Azure Integration Runtime to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to 0.
	// +kubebuilder:validation:Optional
	TimeToLiveMin *float64 `json:"timeToLiveMin,omitempty" tf:"time_to_live_min,omitempty"`
}

// IntegrationRuntimeAzureSpec defines the desired state of IntegrationRuntimeAzure
type IntegrationRuntimeAzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationRuntimeAzureParameters `json:"forProvider"`
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
	InitProvider IntegrationRuntimeAzureInitParameters `json:"initProvider,omitempty"`
}

// IntegrationRuntimeAzureStatus defines the observed state of IntegrationRuntimeAzure.
type IntegrationRuntimeAzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationRuntimeAzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IntegrationRuntimeAzure is the Schema for the IntegrationRuntimeAzures API. Manages a Synapse Azure Integration Runtime.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IntegrationRuntimeAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   IntegrationRuntimeAzureSpec   `json:"spec"`
	Status IntegrationRuntimeAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeAzureList contains a list of IntegrationRuntimeAzures
type IntegrationRuntimeAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationRuntimeAzure `json:"items"`
}

// Repository type metadata.
var (
	IntegrationRuntimeAzure_Kind             = "IntegrationRuntimeAzure"
	IntegrationRuntimeAzure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationRuntimeAzure_Kind}.String()
	IntegrationRuntimeAzure_KindAPIVersion   = IntegrationRuntimeAzure_Kind + "." + CRDGroupVersion.String()
	IntegrationRuntimeAzure_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationRuntimeAzure_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationRuntimeAzure{}, &IntegrationRuntimeAzureList{})
}
