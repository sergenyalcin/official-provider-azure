// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceconfiguration"
)

var maintenanceCrdGroup = "maintenance.azure.upbound.io"

// Setup_maintenance creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_maintenance(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: maintenanceCrdGroup, Kind: "MaintenanceAssignmentDedicatedHost"}:  maintenanceassignmentdedicatedhost.Setup,
		schema.GroupKind{Group: maintenanceCrdGroup, Kind: "MaintenanceAssignmentVirtualMachine"}: maintenanceassignmentvirtualmachine.Setup,
		schema.GroupKind{Group: maintenanceCrdGroup, Kind: "MaintenanceConfiguration"}:            maintenanceconfiguration.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, maintenanceCrdGroup, o); err != nil {
		return err
	}
	return nil
}
