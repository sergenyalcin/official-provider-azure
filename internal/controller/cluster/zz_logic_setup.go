// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	appactioncustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/cluster/logic/appworkflow"
	integrationserviceenvironment "github.com/upbound/provider-azure/internal/controller/cluster/logic/integrationserviceenvironment"
)

var logicCrdGroup = "logic.azure.upbound.io"

// Setup_logic creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logic(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppActionCustom"}:                         appactioncustom.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppActionHTTP"}:                           appactionhttp.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppIntegrationAccount"}:                   appintegrationaccount.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppIntegrationAccountBatchConfiguration"}: appintegrationaccountbatchconfiguration.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppIntegrationAccountPartner"}:            appintegrationaccountpartner.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppIntegrationAccountSchema"}:             appintegrationaccountschema.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppIntegrationAccountSession"}:            appintegrationaccountsession.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppTriggerCustom"}:                        apptriggercustom.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppTriggerHTTPRequest"}:                   apptriggerhttprequest.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppTriggerRecurrence"}:                    apptriggerrecurrence.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "AppWorkflow"}:                             appworkflow.Setup,
		schema.GroupKind{Group: logicCrdGroup, Kind: "IntegrationServiceEnvironment"}:           integrationserviceenvironment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, logicCrdGroup, o); err != nil {
		return err
	}
	return nil
}
