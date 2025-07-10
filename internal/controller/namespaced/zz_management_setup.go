// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	managementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/management/managementgroup"
	managementgroupsubscriptionassociation "github.com/upbound/provider-azure/internal/controller/namespaced/management/managementgroupsubscriptionassociation"
)

var managementCrdGroup = "management.azure.m.upbound.io"

// Setup_management creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_management(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: managementCrdGroup, Kind: "ManagementGroup"}:                        managementgroup.Setup,
		schema.GroupKind{Group: managementCrdGroup, Kind: "ManagementGroupSubscriptionAssociation"}: managementgroupsubscriptionassociation.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, managementCrdGroup, o); err != nil {
		return err
	}
	return nil
}
