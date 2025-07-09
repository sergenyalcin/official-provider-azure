// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/cluster/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/cluster/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/cluster/storage/container"
	containerimmutabilitypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/containerimmutabilitypolicy"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/cluster/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/cluster/storage/objectreplication"
	queue "github.com/upbound/provider-azure/internal/controller/cluster/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/cluster/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/cluster/storage/sharedirectory"
	table "github.com/upbound/provider-azure/internal/controller/cluster/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/cluster/storage/tableentity"
)

var storageCrdGroup = "storage.azure.upbound.io"

// Setup_storage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storage(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: storageCrdGroup, Kind: "Account"}:                     account.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "AccountLocalUser"}:            accountlocaluser.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "AccountNetworkRules"}:         accountnetworkrules.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "Blob"}:                        blob.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "BlobInventoryPolicy"}:         blobinventorypolicy.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "Container"}:                   container.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "ContainerImmutabilityPolicy"}: containerimmutabilitypolicy.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "DataLakeGen2FileSystem"}:      datalakegen2filesystem.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "DataLakeGen2Path"}:            datalakegen2path.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "EncryptionScope"}:             encryptionscope.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "ManagementPolicy"}:            managementpolicy.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "ObjectReplication"}:           objectreplication.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "Queue"}:                       queue.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "Share"}:                       share.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "ShareDirectory"}:              sharedirectory.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "Table"}:                       table.Setup,
		schema.GroupKind{Group: storageCrdGroup, Kind: "TableEntity"}:                 tableentity.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
