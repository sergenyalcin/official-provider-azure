// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicypostgresql"
	backupvault "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/resourceguard"
)

var dataprotectionCrdGroup = "dataprotection.azure.upbound.io"

// Setup_dataprotection creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataprotection(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupInstanceBlobStorage"}:       backupinstanceblobstorage.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupInstanceDisk"}:              backupinstancedisk.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupInstanceKubernetesCluster"}: backupinstancekubernetescluster.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupInstancePostgreSQL"}:        backupinstancepostgresql.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupPolicyBlobStorage"}:         backuppolicyblobstorage.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupPolicyDisk"}:                backuppolicydisk.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupPolicyKubernetesCluster"}:   backuppolicykubernetescluster.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupPolicyPostgreSQL"}:          backuppolicypostgresql.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "BackupVault"}:                     backupvault.Setup,
		schema.GroupKind{Group: dataprotectionCrdGroup, Kind: "ResourceGuard"}:                   resourceguard.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, dataprotectionCrdGroup, o); err != nil {
		return err
	}
	return nil
}
