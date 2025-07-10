// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	cluster "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/job"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputiothub"
)

var streamanalyticsCrdGroup = "streamanalytics.azure.m.upbound.io"

// Setup_streamanalytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_streamanalytics(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "Cluster"}:                cluster.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "FunctionJavascriptUda"}:  functionjavascriptuda.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "Job"}:                    job.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "ManagedPrivateEndpoint"}: managedprivateendpoint.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputBlob"}:             outputblob.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputEventHub"}:         outputeventhub.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputFunction"}:         outputfunction.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputMSSQL"}:            outputmssql.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputPowerBI"}:          outputpowerbi.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputServiceBusQueue"}:  outputservicebusqueue.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputServiceBusTopic"}:  outputservicebustopic.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputSynapse"}:          outputsynapse.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "OutputTable"}:            outputtable.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "ReferenceInputBlob"}:     referenceinputblob.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "ReferenceInputMSSQL"}:    referenceinputmssql.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "StreamInputBlob"}:        streaminputblob.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "StreamInputEventHub"}:    streaminputeventhub.Setup,
		schema.GroupKind{Group: streamanalyticsCrdGroup, Kind: "StreamInputIOTHub"}:      streaminputiothub.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, streamanalyticsCrdGroup, o); err != nil {
		return err
	}
	return nil
}
