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

type SQLDedicatedGatewayInitParameters struct {

	// The instance count for the CosmosDB SQL Dedicated Gateway. Possible value is between 1 and 5.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The instance size for the CosmosDB SQL Dedicated Gateway. Changing this forces a new resource to be created. Possible values are Cosmos.D4s, Cosmos.D8s and Cosmos.D16s.
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`
}

type SQLDedicatedGatewayObservation struct {

	// The resource ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosDBAccountID *string `json:"cosmosdbAccountId,omitempty" tf:"cosmosdb_account_id,omitempty"`

	// The ID of the CosmosDB SQL Dedicated Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance count for the CosmosDB SQL Dedicated Gateway. Possible value is between 1 and 5.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The instance size for the CosmosDB SQL Dedicated Gateway. Changing this forces a new resource to be created. Possible values are Cosmos.D4s, Cosmos.D8s and Cosmos.D16s.
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`
}

type SQLDedicatedGatewayParameters struct {

	// The resource ID of the CosmosDB Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CosmosDBAccountID *string `json:"cosmosdbAccountId,omitempty" tf:"cosmosdb_account_id,omitempty"`

	// Reference to a Account in cosmosdb to populate cosmosdbAccountId.
	// +kubebuilder:validation:Optional
	CosmosDBAccountIDRef *v1.Reference `json:"cosmosdbAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate cosmosdbAccountId.
	// +kubebuilder:validation:Optional
	CosmosDBAccountIDSelector *v1.Selector `json:"cosmosdbAccountIdSelector,omitempty" tf:"-"`

	// The instance count for the CosmosDB SQL Dedicated Gateway. Possible value is between 1 and 5.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The instance size for the CosmosDB SQL Dedicated Gateway. Changing this forces a new resource to be created. Possible values are Cosmos.D4s, Cosmos.D8s and Cosmos.D16s.
	// +kubebuilder:validation:Optional
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`
}

// SQLDedicatedGatewaySpec defines the desired state of SQLDedicatedGateway
type SQLDedicatedGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLDedicatedGatewayParameters `json:"forProvider"`
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
	InitProvider SQLDedicatedGatewayInitParameters `json:"initProvider,omitempty"`
}

// SQLDedicatedGatewayStatus defines the observed state of SQLDedicatedGateway.
type SQLDedicatedGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLDedicatedGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLDedicatedGateway is the Schema for the SQLDedicatedGateways API. Manages a SQL Dedicated Gateway within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLDedicatedGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceCount) || (has(self.initProvider) && has(self.initProvider.instanceCount))",message="spec.forProvider.instanceCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceSize) || (has(self.initProvider) && has(self.initProvider.instanceSize))",message="spec.forProvider.instanceSize is a required parameter"
	Spec   SQLDedicatedGatewaySpec   `json:"spec"`
	Status SQLDedicatedGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDedicatedGatewayList contains a list of SQLDedicatedGateways
type SQLDedicatedGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLDedicatedGateway `json:"items"`
}

// Repository type metadata.
var (
	SQLDedicatedGateway_Kind             = "SQLDedicatedGateway"
	SQLDedicatedGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLDedicatedGateway_Kind}.String()
	SQLDedicatedGateway_KindAPIVersion   = SQLDedicatedGateway_Kind + "." + CRDGroupVersion.String()
	SQLDedicatedGateway_GroupVersionKind = CRDGroupVersion.WithKind(SQLDedicatedGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLDedicatedGateway{}, &SQLDedicatedGatewayList{})
}
