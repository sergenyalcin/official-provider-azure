// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	diskpool "github.com/upbound/provider-azure/internal/controller/cluster/storagepool/diskpool"
)

var storagepoolCrdGroup = "storagepool.azure.upbound.io"

// Setup_storagepool creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storagepool(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: storagepoolCrdGroup, Kind: "DiskPool"}: diskpool.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, storagepoolCrdGroup, o); err != nil {
		return err
	}
	return nil
}
