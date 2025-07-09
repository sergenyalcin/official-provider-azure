// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryreplicationpolicy"
	vault "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/vault"
)

var recoveryservicesCrdGroup = "recoveryservices.azure.m.upbound.io"

// Setup_recoveryservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_recoveryservices(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupContainerStorageAccount"}:          backupcontainerstorageaccount.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupPolicyFileShare"}:                  backuppolicyfileshare.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupPolicyVM"}:                         backuppolicyvm.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupPolicyVMWorkload"}:                 backuppolicyvmworkload.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupProtectedFileShare"}:               backupprotectedfileshare.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "BackupProtectedVM"}:                      backupprotectedvm.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "SiteRecoveryFabric"}:                     siterecoveryfabric.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "SiteRecoveryNetworkMapping"}:             siterecoverynetworkmapping.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "SiteRecoveryProtectionContainer"}:        siterecoveryprotectioncontainer.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "SiteRecoveryProtectionContainerMapping"}: siterecoveryprotectioncontainermapping.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "SiteRecoveryReplicationPolicy"}:          siterecoveryreplicationpolicy.Setup,
		schema.GroupKind{Group: recoveryservicesCrdGroup, Kind: "Vault"}:                                  vault.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
