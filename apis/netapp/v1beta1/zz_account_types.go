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

type AccountInitParameters struct {

	// A active_directory block as defined below.
	ActiveDirectory []ActiveDirectoryInitParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// The identity block where it is used when customer managed keys based encryption will be enabled as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountObservation struct {

	// A active_directory block as defined below.
	ActiveDirectory []ActiveDirectoryObservation `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// The ID of the NetApp Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identity block where it is used when customer managed keys based encryption will be enabled as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountParameters struct {

	// A active_directory block as defined below.
	// +kubebuilder:validation:Optional
	ActiveDirectory []ActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// The identity block where it is used when customer managed keys based encryption will be enabled as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
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
}

type ActiveDirectoryInitParameters struct {

	// If enabled, AES encryption will be enabled for SMB communication. Defaults to false.
	// If enabled, AES encryption will be enabled for SMB communication.
	AesEncryptionEnabled *bool `json:"aesEncryptionEnabled,omitempty" tf:"aes_encryption_enabled,omitempty"`

	// A list of DNS server IP addresses for the Active Directory domain. Only allows IPv4 address.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The name of the Active Directory domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Name of the active directory machine.
	// Name of the active directory machine. This optional parameter is used only while creating kerberos volume.
	KerberosAdName *string `json:"kerberosAdName,omitempty" tf:"kerberos_ad_name,omitempty"`

	// kdc server IP addresses for the active directory machine.
	// IP address of the KDC server (usually same the DC). This optional parameter is used only while creating kerberos volume.
	KerberosKdcIP *string `json:"kerberosKdcIp,omitempty" tf:"kerberos_kdc_ip,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be secured via TLS. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be secured via TLS.
	LdapOverTLSEnabled *bool `json:"ldapOverTlsEnabled,omitempty" tf:"ldap_over_tls_enabled,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be signed. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be signed.
	LdapSigningEnabled *bool `json:"ldapSigningEnabled,omitempty" tf:"ldap_signing_enabled,omitempty"`

	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes. Defaults to false.
	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.
	LocalNFSUsersWithLdapAllowed *bool `json:"localNfsUsersWithLdapAllowed,omitempty" tf:"local_nfs_users_with_ldap_allowed,omitempty"`

	// The Organizational Unit (OU) within Active Directory where machines will be created. If blank, defaults to CN=Computers.
	// The Organizational Unit (OU) within the Windows Active Directory where machines will be created. If blank, defaults to 'CN=Computers'
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
	SMBServerName *string `json:"smbServerName,omitempty" tf:"smb_server_name,omitempty"`

	// When LDAP over SSL/TLS is enabled, the LDAP client is required to have a base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes. Required if ldap_over_tls_enabled is set to true.
	// When LDAP over SSL/TLS is enabled, the LDAP client is required to have base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes.
	ServerRootCACertificateSecretRef *v1.SecretKeySelector `json:"serverRootCaCertificateSecretRef,omitempty" tf:"-"`

	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to Default-First-Site-Name.
	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to 'Default-First-Site-Name'
	SiteName *string `json:"siteName,omitempty" tf:"site_name,omitempty"`

	// The Username of Active Directory Domain Administrator.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ActiveDirectoryObservation struct {

	// If enabled, AES encryption will be enabled for SMB communication. Defaults to false.
	// If enabled, AES encryption will be enabled for SMB communication.
	AesEncryptionEnabled *bool `json:"aesEncryptionEnabled,omitempty" tf:"aes_encryption_enabled,omitempty"`

	// A list of DNS server IP addresses for the Active Directory domain. Only allows IPv4 address.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The name of the Active Directory domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Name of the active directory machine.
	// Name of the active directory machine. This optional parameter is used only while creating kerberos volume.
	KerberosAdName *string `json:"kerberosAdName,omitempty" tf:"kerberos_ad_name,omitempty"`

	// kdc server IP addresses for the active directory machine.
	// IP address of the KDC server (usually same the DC). This optional parameter is used only while creating kerberos volume.
	KerberosKdcIP *string `json:"kerberosKdcIp,omitempty" tf:"kerberos_kdc_ip,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be secured via TLS. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be secured via TLS.
	LdapOverTLSEnabled *bool `json:"ldapOverTlsEnabled,omitempty" tf:"ldap_over_tls_enabled,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be signed. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be signed.
	LdapSigningEnabled *bool `json:"ldapSigningEnabled,omitempty" tf:"ldap_signing_enabled,omitempty"`

	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes. Defaults to false.
	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.
	LocalNFSUsersWithLdapAllowed *bool `json:"localNfsUsersWithLdapAllowed,omitempty" tf:"local_nfs_users_with_ldap_allowed,omitempty"`

	// The Organizational Unit (OU) within Active Directory where machines will be created. If blank, defaults to CN=Computers.
	// The Organizational Unit (OU) within the Windows Active Directory where machines will be created. If blank, defaults to 'CN=Computers'
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
	SMBServerName *string `json:"smbServerName,omitempty" tf:"smb_server_name,omitempty"`

	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to Default-First-Site-Name.
	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to 'Default-First-Site-Name'
	SiteName *string `json:"siteName,omitempty" tf:"site_name,omitempty"`

	// The Username of Active Directory Domain Administrator.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ActiveDirectoryParameters struct {

	// If enabled, AES encryption will be enabled for SMB communication. Defaults to false.
	// If enabled, AES encryption will be enabled for SMB communication.
	// +kubebuilder:validation:Optional
	AesEncryptionEnabled *bool `json:"aesEncryptionEnabled,omitempty" tf:"aes_encryption_enabled,omitempty"`

	// A list of DNS server IP addresses for the Active Directory domain. Only allows IPv4 address.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers" tf:"dns_servers,omitempty"`

	// The name of the Active Directory domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Name of the active directory machine.
	// Name of the active directory machine. This optional parameter is used only while creating kerberos volume.
	// +kubebuilder:validation:Optional
	KerberosAdName *string `json:"kerberosAdName,omitempty" tf:"kerberos_ad_name,omitempty"`

	// kdc server IP addresses for the active directory machine.
	// IP address of the KDC server (usually same the DC). This optional parameter is used only while creating kerberos volume.
	// +kubebuilder:validation:Optional
	KerberosKdcIP *string `json:"kerberosKdcIp,omitempty" tf:"kerberos_kdc_ip,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be secured via TLS. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be secured via TLS.
	// +kubebuilder:validation:Optional
	LdapOverTLSEnabled *bool `json:"ldapOverTlsEnabled,omitempty" tf:"ldap_over_tls_enabled,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be signed. Defaults to false.
	// Specifies whether or not the LDAP traffic needs to be signed.
	// +kubebuilder:validation:Optional
	LdapSigningEnabled *bool `json:"ldapSigningEnabled,omitempty" tf:"ldap_signing_enabled,omitempty"`

	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes. Defaults to false.
	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.
	// +kubebuilder:validation:Optional
	LocalNFSUsersWithLdapAllowed *bool `json:"localNfsUsersWithLdapAllowed,omitempty" tf:"local_nfs_users_with_ldap_allowed,omitempty"`

	// The Organizational Unit (OU) within Active Directory where machines will be created. If blank, defaults to CN=Computers.
	// The Organizational Unit (OU) within the Windows Active Directory where machines will be created. If blank, defaults to 'CN=Computers'
	// +kubebuilder:validation:Optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// The password associated with the username.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
	// +kubebuilder:validation:Optional
	SMBServerName *string `json:"smbServerName" tf:"smb_server_name,omitempty"`

	// When LDAP over SSL/TLS is enabled, the LDAP client is required to have a base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes. Required if ldap_over_tls_enabled is set to true.
	// When LDAP over SSL/TLS is enabled, the LDAP client is required to have base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes.
	// +kubebuilder:validation:Optional
	ServerRootCACertificateSecretRef *v1.SecretKeySelector `json:"serverRootCaCertificateSecretRef,omitempty" tf:"-"`

	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to Default-First-Site-Name.
	// The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to 'Default-First-Site-Name'
	// +kubebuilder:validation:Optional
	SiteName *string `json:"siteName,omitempty" tf:"site_name,omitempty"`

	// The Username of Active Directory Domain Administrator.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type IdentityInitParameters struct {

	// The identity id of the user assigned identity to use when type is UserAssigned
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The identity type, which can be SystemAssigned or UserAssigned. Only one type at a time is supported by Azure NetApp Files.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// The identity id of the user assigned identity to use when type is UserAssigned
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The ID of the NetApp Account.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the NetApp Account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The identity type, which can be SystemAssigned or UserAssigned. Only one type at a time is supported by Azure NetApp Files.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// The identity id of the user assigned identity to use when type is UserAssigned
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The identity type, which can be SystemAssigned or UserAssigned. Only one type at a time is supported by Azure NetApp Files.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
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
	InitProvider AccountInitParameters `json:"initProvider,omitempty"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Account is the Schema for the Accounts API. Manages a NetApp Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
