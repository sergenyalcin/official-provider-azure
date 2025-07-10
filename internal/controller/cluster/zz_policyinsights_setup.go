// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/subscriptionpolicyremediation"
)

var policyinsightsCrdGroup = "policyinsights.azure.upbound.io"

// Setup_policyinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_policyinsights(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: policyinsightsCrdGroup, Kind: "ResourcePolicyRemediation"}:     resourcepolicyremediation.Setup,
		schema.GroupKind{Group: policyinsightsCrdGroup, Kind: "SubscriptionPolicyRemediation"}: subscriptionpolicyremediation.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, policyinsightsCrdGroup, o); err != nil {
		return err
	}
	return nil
}
