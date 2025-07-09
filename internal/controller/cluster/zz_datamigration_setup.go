// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationservice"
)

var datamigrationCrdGroup = "datamigration.azure.upbound.io"

// Setup_datamigration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datamigration(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: datamigrationCrdGroup, Kind: "DatabaseMigrationProject"}: databasemigrationproject.Setup,
		schema.GroupKind{Group: datamigrationCrdGroup, Kind: "DatabaseMigrationService"}: databasemigrationservice.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
