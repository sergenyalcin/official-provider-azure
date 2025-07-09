// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	accesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/accesspolicy"
	certificate "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/vault"
)

var keyvaultCrdGroup = "keyvault.azure.m.upbound.io"

// Setup_keyvault creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_keyvault(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "AccessPolicy"}:                            accesspolicy.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "Certificate"}:                             certificate.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "CertificateContacts"}:                     certificatecontacts.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "CertificateIssuer"}:                       certificateissuer.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "Key"}:                                     key.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "ManagedHardwareSecurityModule"}:           managedhardwaresecuritymodule.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "ManagedStorageAccount"}:                   managedstorageaccount.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "ManagedStorageAccountSASTokenDefinition"}: managedstorageaccountsastokendefinition.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "Secret"}:                                  secret.Setup,
		schema.GroupKind{Group: keyvaultCrdGroup, Kind: "Vault"}:                                   vault.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
