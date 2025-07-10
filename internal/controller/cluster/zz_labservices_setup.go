// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	labservicelab "github.com/upbound/provider-azure/internal/controller/cluster/labservices/labservicelab"
	labserviceplan "github.com/upbound/provider-azure/internal/controller/cluster/labservices/labserviceplan"
)

var labservicesCrdGroup = "labservices.azure.upbound.io"

// Setup_labservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_labservices(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: labservicesCrdGroup, Kind: "LabServiceLab"}:  labservicelab.Setup,
		schema.GroupKind{Group: labservicesCrdGroup, Kind: "LabServicePlan"}: labserviceplan.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, labservicesCrdGroup, o); err != nil {
		return err
	}
	return nil
}
