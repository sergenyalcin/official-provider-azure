// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	healthcaredicomservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareworkspace"
)

var healthcareapisCrdGroup = "healthcareapis.azure.upbound.io"

// Setup_healthcareapis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthcareapis(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareDICOMService"}:                  healthcaredicomservice.Setup,
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareFHIRService"}:                   healthcarefhirservice.Setup,
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareMedtechService"}:                healthcaremedtechservice.Setup,
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareMedtechServiceFHIRDestination"}: healthcaremedtechservicefhirdestination.Setup,
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareService"}:                       healthcareservice.Setup,
		schema.GroupKind{Group: healthcareapisCrdGroup, Kind: "HealthcareWorkspace"}:                     healthcareworkspace.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, healthcareapisCrdGroup, o); err != nil {
		return err
	}
	return nil
}
