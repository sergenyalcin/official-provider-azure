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

type BackupPolicyBlobStorageInitParameters struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval. Changing this forces a new Backup Policy Blob Storage to be created.
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The duration of operational default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	OperationalDefaultRetentionDuration *string `json:"operationalDefaultRetentionDuration,omitempty" tf:"operational_default_retention_duration,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionRule []BackupPolicyBlobStorageRetentionRuleInitParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// Specifies the Time Zone which should be used by the backup schedule. Changing this forces a new Backup Policy Blob Storage to be created.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// The duration of vault default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultDefaultRetentionDuration *string `json:"vaultDefaultRetentionDuration,omitempty" tf:"vault_default_retention_duration,omitempty"`
}

type BackupPolicyBlobStorageObservation struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval. Changing this forces a new Backup Policy Blob Storage to be created.
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The ID of the Backup Policy Blob Storage.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The duration of operational default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	OperationalDefaultRetentionDuration *string `json:"operationalDefaultRetentionDuration,omitempty" tf:"operational_default_retention_duration,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionRule []BackupPolicyBlobStorageRetentionRuleObservation `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// Specifies the Time Zone which should be used by the backup schedule. Changing this forces a new Backup Policy Blob Storage to be created.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// The duration of vault default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultDefaultRetentionDuration *string `json:"vaultDefaultRetentionDuration,omitempty" tf:"vault_default_retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupPolicyBlobStorageParameters struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The duration of operational default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	OperationalDefaultRetentionDuration *string `json:"operationalDefaultRetentionDuration,omitempty" tf:"operational_default_retention_duration,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	RetentionRule []BackupPolicyBlobStorageRetentionRuleParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// Specifies the Time Zone which should be used by the backup schedule. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// The duration of vault default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	VaultDefaultRetentionDuration *string `json:"vaultDefaultRetentionDuration,omitempty" tf:"vault_default_retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta2.BackupVault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

type BackupPolicyBlobStorageRetentionRuleCriteriaInitParameters struct {

	// Possible values are AllBackup, FirstOfDay, FirstOfWeek, FirstOfMonth and FirstOfYear. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy Blob Storage to be created.
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`

	// Must be between 0 and 28. 0 for last day within the month. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Possible values are Monday, Tuesday, Thursday, Friday, Saturday and Sunday. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Possible values are January, February, March, April, May, June, July, August, September, October, November and December. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	MonthsOfYear []*string `json:"monthsOfYear,omitempty" tf:"months_of_year,omitempty"`

	// Specifies a list of backup times for backup in the RFC3339 format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	ScheduledBackupTimes []*string `json:"scheduledBackupTimes,omitempty" tf:"scheduled_backup_times,omitempty"`

	// Possible values are First, Second, Third, Fourth and Last. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	WeeksOfMonth []*string `json:"weeksOfMonth,omitempty" tf:"weeks_of_month,omitempty"`
}

type BackupPolicyBlobStorageRetentionRuleCriteriaObservation struct {

	// Possible values are AllBackup, FirstOfDay, FirstOfWeek, FirstOfMonth and FirstOfYear. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy Blob Storage to be created.
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`

	// Must be between 0 and 28. 0 for last day within the month. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Possible values are Monday, Tuesday, Thursday, Friday, Saturday and Sunday. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Possible values are January, February, March, April, May, June, July, August, September, October, November and December. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	MonthsOfYear []*string `json:"monthsOfYear,omitempty" tf:"months_of_year,omitempty"`

	// Specifies a list of backup times for backup in the RFC3339 format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	ScheduledBackupTimes []*string `json:"scheduledBackupTimes,omitempty" tf:"scheduled_backup_times,omitempty"`

	// Possible values are First, Second, Third, Fourth and Last. Changing this forces a new Backup Policy Blob Storage to be created.
	// +listType=set
	WeeksOfMonth []*string `json:"weeksOfMonth,omitempty" tf:"weeks_of_month,omitempty"`
}

type BackupPolicyBlobStorageRetentionRuleCriteriaParameters struct {

	// Possible values are AllBackup, FirstOfDay, FirstOfWeek, FirstOfMonth and FirstOfYear. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`

	// Must be between 0 and 28. 0 for last day within the month. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Possible values are Monday, Tuesday, Thursday, Friday, Saturday and Sunday. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Possible values are January, February, March, April, May, June, July, August, September, October, November and December. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	MonthsOfYear []*string `json:"monthsOfYear,omitempty" tf:"months_of_year,omitempty"`

	// Specifies a list of backup times for backup in the RFC3339 format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	ScheduledBackupTimes []*string `json:"scheduledBackupTimes,omitempty" tf:"scheduled_backup_times,omitempty"`

	// Possible values are First, Second, Third, Fourth and Last. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	WeeksOfMonth []*string `json:"weeksOfMonth,omitempty" tf:"weeks_of_month,omitempty"`
}

type BackupPolicyBlobStorageRetentionRuleInitParameters struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	Criteria *BackupPolicyBlobStorageRetentionRuleCriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// A life_cycle block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	LifeCycle *LifeCycleInitParameters `json:"lifeCycle,omitempty" tf:"life_cycle,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Blob Storage to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Blob Storage to be created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type BackupPolicyBlobStorageRetentionRuleObservation struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	Criteria *BackupPolicyBlobStorageRetentionRuleCriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// A life_cycle block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	LifeCycle *LifeCycleObservation `json:"lifeCycle,omitempty" tf:"life_cycle,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Blob Storage to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Blob Storage to be created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type BackupPolicyBlobStorageRetentionRuleParameters struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	Criteria *BackupPolicyBlobStorageRetentionRuleCriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// A life_cycle block as defined below. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	LifeCycle *LifeCycleParameters `json:"lifeCycle" tf:"life_cycle,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

type LifeCycleInitParameters struct {

	// The type of data store. The only possible value is VaultStore. Changing this forces a new Backup Policy Blob Storage to be created.
	DataStoreType *string `json:"dataStoreType,omitempty" tf:"data_store_type,omitempty"`

	// Duration after which the backup is deleted. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type LifeCycleObservation struct {

	// The type of data store. The only possible value is VaultStore. Changing this forces a new Backup Policy Blob Storage to be created.
	DataStoreType *string `json:"dataStoreType,omitempty" tf:"data_store_type,omitempty"`

	// Duration after which the backup is deleted. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type LifeCycleParameters struct {

	// The type of data store. The only possible value is VaultStore. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	DataStoreType *string `json:"dataStoreType" tf:"data_store_type,omitempty"`

	// Duration after which the backup is deleted. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`
}

// BackupPolicyBlobStorageSpec defines the desired state of BackupPolicyBlobStorage
type BackupPolicyBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyBlobStorageParameters `json:"forProvider"`
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
	InitProvider BackupPolicyBlobStorageInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyBlobStorageStatus defines the observed state of BackupPolicyBlobStorage.
type BackupPolicyBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupPolicyBlobStorage is the Schema for the BackupPolicyBlobStorages API. Manages a Backup Policy Blob Storage.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicyBlobStorageSpec   `json:"spec"`
	Status            BackupPolicyBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyBlobStorageList contains a list of BackupPolicyBlobStorages
type BackupPolicyBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyBlobStorage_Kind             = "BackupPolicyBlobStorage"
	BackupPolicyBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyBlobStorage_Kind}.String()
	BackupPolicyBlobStorage_KindAPIVersion   = BackupPolicyBlobStorage_Kind + "." + CRDGroupVersion.String()
	BackupPolicyBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyBlobStorage{}, &BackupPolicyBlobStorageList{})
}
