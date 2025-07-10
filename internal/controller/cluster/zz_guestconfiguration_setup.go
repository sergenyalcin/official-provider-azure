// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	policyvirtualmachineconfigurationassignment "github.com/upbound/provider-azure/internal/controller/cluster/guestconfiguration/policyvirtualmachineconfigurationassignment"
)

var guestconfigurationCrdGroup = "guestconfiguration.azure.upbound.io"

// Setup_guestconfiguration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_guestconfiguration(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: guestconfigurationCrdGroup, Kind: "PolicyVirtualMachineConfigurationAssignment"}: policyvirtualmachineconfigurationassignment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, guestconfigurationCrdGroup, o); err != nil {
		return err
	}
	return nil
}
