// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/namespaced/automation/account"
	connection "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/namespaced/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/namespaced/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/namespaced/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/namespaced/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/namespaced/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/namespaced/automation/webhook"
)

var automationCrdGroup = "automation.azure.m.upbound.io"

// Setup_automation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_automation(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: automationCrdGroup, Kind: "Account"}:                      account.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "Connection"}:                   connection.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "ConnectionClassicCertificate"}: connectionclassiccertificate.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "ConnectionType"}:               connectiontype.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "Credential"}:                   credential.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "HybridRunBookWorkerGroup"}:     hybridrunbookworkergroup.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "Module"}:                       module.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "RunBook"}:                      runbook.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "Schedule"}:                     schedule.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "VariableBool"}:                 variablebool.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "VariableDateTime"}:             variabledatetime.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "VariableInt"}:                  variableint.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "VariableString"}:               variablestring.Setup,
		schema.GroupKind{Group: automationCrdGroup, Kind: "Webhook"}:                      webhook.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, automationCrdGroup, o); err != nil {
		return err
	}
	return nil
}
