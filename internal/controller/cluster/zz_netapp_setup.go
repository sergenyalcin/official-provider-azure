// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/cluster/netapp/account"
	pool "github.com/upbound/provider-azure/internal/controller/cluster/netapp/pool"
	snapshot "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/cluster/netapp/volume"
)

var netappCrdGroup = "netapp.azure.upbound.io"

// Setup_netapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_netapp(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: netappCrdGroup, Kind: "Account"}:        account.Setup,
		schema.GroupKind{Group: netappCrdGroup, Kind: "Pool"}:           pool.Setup,
		schema.GroupKind{Group: netappCrdGroup, Kind: "Snapshot"}:       snapshot.Setup,
		schema.GroupKind{Group: netappCrdGroup, Kind: "SnapshotPolicy"}: snapshotpolicy.Setup,
		schema.GroupKind{Group: netappCrdGroup, Kind: "Volume"}:         volume.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, netappCrdGroup, o); err != nil {
		return err
	}
	return nil
}
