// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	containerapp "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/containerapp"
	customdomain "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentstorage"
)

var containerappCrdGroup = "containerapp.azure.m.upbound.io"

// Setup_containerapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerapp(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: containerappCrdGroup, Kind: "ContainerApp"}:             containerapp.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "CustomDomain"}:             customdomain.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "Environment"}:              environment.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "EnvironmentCertificate"}:   environmentcertificate.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "EnvironmentCustomDomain"}:  environmentcustomdomain.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "EnvironmentDaprComponent"}: environmentdaprcomponent.Setup,
		schema.GroupKind{Group: containerappCrdGroup, Kind: "EnvironmentStorage"}:       environmentstorage.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, containerappCrdGroup, o); err != nil {
		return err
	}
	return nil
}
