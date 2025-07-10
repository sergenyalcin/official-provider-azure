// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	storagesync "github.com/upbound/provider-azure/internal/controller/cluster/storagesync/storagesync"
)

var storagesyncCrdGroup = "storagesync.azure.upbound.io"

// Setup_storagesync creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storagesync(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: storagesyncCrdGroup, Kind: "StorageSync"}: storagesync.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, storagesyncCrdGroup, o); err != nil {
		return err
	}
	return nil
}
