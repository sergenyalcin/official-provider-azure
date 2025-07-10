// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryreplicationpolicy"
	vault "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/vault"
)

var recoveryservicesCrdGroup = "recoveryservices.azure.upbound.io"

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
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, recoveryservicesCrdGroup, o); err != nil {
		return err
	}
	return nil
}
