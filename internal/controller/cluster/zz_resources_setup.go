// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/subscriptiontemplatedeployment"
)

var resourcesCrdGroup = "resources.azure.upbound.io"

// Setup_resources creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_resources(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: resourcesCrdGroup, Kind: "ResourceDeploymentScriptAzureCli"}:        resourcedeploymentscriptazurecli.Setup,
		schema.GroupKind{Group: resourcesCrdGroup, Kind: "ResourceDeploymentScriptAzurePowerShell"}: resourcedeploymentscriptazurepowershell.Setup,
		schema.GroupKind{Group: resourcesCrdGroup, Kind: "ResourceGroupTemplateDeployment"}:         resourcegrouptemplatedeployment.Setup,
		schema.GroupKind{Group: resourcesCrdGroup, Kind: "SubscriptionTemplateDeployment"}:          subscriptiontemplatedeployment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
