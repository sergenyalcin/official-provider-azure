// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	accessconnector "github.com/upbound/provider-azure/internal/controller/cluster/databricks/accessconnector"
	workspace "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspace"
	workspacecustomermanagedkey "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspacecustomermanagedkey"
	workspacerootdbfscustomermanagedkey "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspacerootdbfscustomermanagedkey"
)

var databricksCrdGroup = "databricks.azure.upbound.io"

// Setup_databricks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_databricks(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: databricksCrdGroup, Kind: "AccessConnector"}:                     accessconnector.Setup,
		schema.GroupKind{Group: databricksCrdGroup, Kind: "Workspace"}:                           workspace.Setup,
		schema.GroupKind{Group: databricksCrdGroup, Kind: "WorkspaceCustomerManagedKey"}:         workspacecustomermanagedkey.Setup,
		schema.GroupKind{Group: databricksCrdGroup, Kind: "WorkspaceRootDbfsCustomerManagedKey"}: workspacerootdbfscustomermanagedkey.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
