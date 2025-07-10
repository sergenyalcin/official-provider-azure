// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	budgetmanagementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetsubscription"
)

var consumptionCrdGroup = "consumption.azure.m.upbound.io"

// Setup_consumption creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_consumption(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: consumptionCrdGroup, Kind: "BudgetManagementGroup"}: budgetmanagementgroup.Setup,
		schema.GroupKind{Group: consumptionCrdGroup, Kind: "BudgetResourceGroup"}:   budgetresourcegroup.Setup,
		schema.GroupKind{Group: consumptionCrdGroup, Kind: "BudgetSubscription"}:    budgetsubscription.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, consumptionCrdGroup, o); err != nil {
		return err
	}
	return nil
}
