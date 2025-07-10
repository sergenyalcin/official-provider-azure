// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	appactiveslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/namespaced/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/namespaced/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/namespaced/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/namespaced/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/namespaced/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowswebappslot"
)

var webCrdGroup = "web.azure.m.upbound.io"

// Setup_web creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_web(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: webCrdGroup, Kind: "AppActiveSlot"}:               appactiveslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "AppHybridConnection"}:         apphybridconnection.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "AppServicePlan"}:              appserviceplan.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "FunctionApp"}:                 functionapp.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "FunctionAppActiveSlot"}:       functionappactiveslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "FunctionAppFunction"}:         functionappfunction.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "FunctionAppHybridConnection"}: functionapphybridconnection.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "FunctionAppSlot"}:             functionappslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "LinuxFunctionApp"}:            linuxfunctionapp.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "LinuxFunctionAppSlot"}:        linuxfunctionappslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "LinuxWebApp"}:                 linuxwebapp.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "LinuxWebAppSlot"}:             linuxwebappslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "ServicePlan"}:                 serviceplan.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "SourceControlToken"}:          sourcecontroltoken.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "StaticSite"}:                  staticsite.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "WindowsFunctionApp"}:          windowsfunctionapp.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "WindowsFunctionAppSlot"}:      windowsfunctionappslot.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "WindowsWebApp"}:               windowswebapp.Setup,
		schema.GroupKind{Group: webCrdGroup, Kind: "WindowsWebAppSlot"}:           windowswebappslot.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, webCrdGroup, o); err != nil {
		return err
	}
	return nil
}
