// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	resourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourceproviderregistration"
	subscription "github.com/upbound/provider-azure/internal/controller/cluster/azure/subscription"
)

var azureCrdGroup = "azure.azure.upbound.io"

// Setup_azure creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_azure(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: azureCrdGroup, Kind: "ResourceGroup"}:                resourcegroup.Setup,
		schema.GroupKind{Group: azureCrdGroup, Kind: "ResourceProviderRegistration"}: resourceproviderregistration.Setup,
		schema.GroupKind{Group: azureCrdGroup, Kind: "Subscription"}:                 subscription.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, azureCrdGroup, o); err != nil {
		return err
	}
	return nil
}
