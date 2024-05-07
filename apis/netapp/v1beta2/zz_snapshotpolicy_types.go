// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DailyScheduleInitParameters struct {

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type DailyScheduleObservation struct {

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type DailyScheduleParameters struct {

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	// +kubebuilder:validation:Optional
	Minute *float64 `json:"minute" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	// +kubebuilder:validation:Optional
	SnapshotsToKeep *float64 `json:"snapshotsToKeep" tf:"snapshots_to_keep,omitempty"`
}

type HourlyScheduleInitParameters struct {

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type HourlyScheduleObservation struct {

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type HourlyScheduleParameters struct {

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	// +kubebuilder:validation:Optional
	Minute *float64 `json:"minute" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	// +kubebuilder:validation:Optional
	SnapshotsToKeep *float64 `json:"snapshotsToKeep" tf:"snapshots_to_keep,omitempty"`
}

type MonthlyScheduleInitParameters struct {

	// List of the days of the month when the snapshots will be created, valid range is from 1 to 30.
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type MonthlyScheduleObservation struct {

	// List of the days of the month when the snapshots will be created, valid range is from 1 to 30.
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type MonthlyScheduleParameters struct {

	// List of the days of the month when the snapshots will be created, valid range is from 1 to 30.
	// +kubebuilder:validation:Optional
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth" tf:"days_of_month,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	// +kubebuilder:validation:Optional
	Minute *float64 `json:"minute" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	// +kubebuilder:validation:Optional
	SnapshotsToKeep *float64 `json:"snapshotsToKeep" tf:"snapshots_to_keep,omitempty"`
}

type SnapshotPolicyInitParameters struct {

	// Sets a daily snapshot schedule. A daily_schedule block as defined below.
	DailySchedule *DailyScheduleInitParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// Defines that the NetApp Snapshot Policy is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Sets an hourly snapshot schedule. A hourly_schedule block as defined below.
	HourlySchedule *HourlyScheduleInitParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets a monthly snapshot schedule. A monthly_schedule block as defined below.
	MonthlySchedule *MonthlyScheduleInitParameters `json:"monthlySchedule,omitempty" tf:"monthly_schedule,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets a weekly snapshot schedule. A weekly_schedule block as defined below.
	WeeklySchedule *WeeklyScheduleInitParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type SnapshotPolicyObservation struct {

	// The name of the NetApp Account in which the NetApp Snapshot Policy should be created. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Sets a daily snapshot schedule. A daily_schedule block as defined below.
	DailySchedule *DailyScheduleObservation `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// Defines that the NetApp Snapshot Policy is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Sets an hourly snapshot schedule. A hourly_schedule block as defined below.
	HourlySchedule *HourlyScheduleObservation `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// The ID of the NetApp Snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets a monthly snapshot schedule. A monthly_schedule block as defined below.
	MonthlySchedule *MonthlyScheduleObservation `json:"monthlySchedule,omitempty" tf:"monthly_schedule,omitempty"`

	// The name of the resource group where the NetApp Snapshot Policy should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets a weekly snapshot schedule. A weekly_schedule block as defined below.
	WeeklySchedule *WeeklyScheduleObservation `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type SnapshotPolicyParameters struct {

	// The name of the NetApp Account in which the NetApp Snapshot Policy should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// Sets a daily snapshot schedule. A daily_schedule block as defined below.
	// +kubebuilder:validation:Optional
	DailySchedule *DailyScheduleParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// Defines that the NetApp Snapshot Policy is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Sets an hourly snapshot schedule. A hourly_schedule block as defined below.
	// +kubebuilder:validation:Optional
	HourlySchedule *HourlyScheduleParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets a monthly snapshot schedule. A monthly_schedule block as defined below.
	// +kubebuilder:validation:Optional
	MonthlySchedule *MonthlyScheduleParameters `json:"monthlySchedule,omitempty" tf:"monthly_schedule,omitempty"`

	// The name of the resource group where the NetApp Snapshot Policy should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets a weekly snapshot schedule. A weekly_schedule block as defined below.
	// +kubebuilder:validation:Optional
	WeeklySchedule *WeeklyScheduleParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type WeeklyScheduleInitParameters struct {

	// List of the week days using English names when the snapshots will be created.
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type WeeklyScheduleObservation struct {

	// List of the week days using English names when the snapshots will be created.
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	Minute *float64 `json:"minute,omitempty" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty" tf:"snapshots_to_keep,omitempty"`
}

type WeeklyScheduleParameters struct {

	// List of the week days using English names when the snapshots will be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek" tf:"days_of_week,omitempty"`

	// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour" tf:"hour,omitempty"`

	// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
	// +kubebuilder:validation:Optional
	Minute *float64 `json:"minute" tf:"minute,omitempty"`

	// How many hourly snapshots to keep, valid range is from 0 to 255.
	// +kubebuilder:validation:Optional
	SnapshotsToKeep *float64 `json:"snapshotsToKeep" tf:"snapshots_to_keep,omitempty"`
}

// SnapshotPolicySpec defines the desired state of SnapshotPolicy
type SnapshotPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotPolicyParameters `json:"forProvider"`
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
	InitProvider SnapshotPolicyInitParameters `json:"initProvider,omitempty"`
}

// SnapshotPolicyStatus defines the observed state of SnapshotPolicy.
type SnapshotPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SnapshotPolicy is the Schema for the SnapshotPolicys API. Manages a NetApp Snapshot Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SnapshotPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   SnapshotPolicySpec   `json:"spec"`
	Status SnapshotPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotPolicyList contains a list of SnapshotPolicys
type SnapshotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotPolicy `json:"items"`
}

// Repository type metadata.
var (
	SnapshotPolicy_Kind             = "SnapshotPolicy"
	SnapshotPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotPolicy_Kind}.String()
	SnapshotPolicy_KindAPIVersion   = SnapshotPolicy_Kind + "." + CRDGroupVersion.String()
	SnapshotPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotPolicy{}, &SnapshotPolicyList{})
}
