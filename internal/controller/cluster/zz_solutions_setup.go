// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/cluster/solutions/managedapplicationdefinition"
)

var solutionsCrdGroup = "solutions.azure.upbound.io"

// Setup_solutions creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_solutions(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: solutionsCrdGroup, Kind: "ManagedApplicationDefinition"}: managedapplicationdefinition.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
