// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	eventsourceeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/timeseriesinsights/eventsourceeventhub"
	eventsourceiothub "github.com/upbound/provider-azure/internal/controller/namespaced/timeseriesinsights/eventsourceiothub"
	gen2environment "github.com/upbound/provider-azure/internal/controller/namespaced/timeseriesinsights/gen2environment"
	referencedataset "github.com/upbound/provider-azure/internal/controller/namespaced/timeseriesinsights/referencedataset"
	standardenvironment "github.com/upbound/provider-azure/internal/controller/namespaced/timeseriesinsights/standardenvironment"
)

var timeseriesinsightsCrdGroup = "timeseriesinsights.azure.m.upbound.io"

// Setup_timeseriesinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_timeseriesinsights(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: timeseriesinsightsCrdGroup, Kind: "EventSourceEventHub"}: eventsourceeventhub.Setup,
		schema.GroupKind{Group: timeseriesinsightsCrdGroup, Kind: "EventSourceIOTHub"}:   eventsourceiothub.Setup,
		schema.GroupKind{Group: timeseriesinsightsCrdGroup, Kind: "Gen2Environment"}:     gen2environment.Setup,
		schema.GroupKind{Group: timeseriesinsightsCrdGroup, Kind: "ReferenceDataSet"}:    referencedataset.Setup,
		schema.GroupKind{Group: timeseriesinsightsCrdGroup, Kind: "StandardEnvironment"}: standardenvironment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
