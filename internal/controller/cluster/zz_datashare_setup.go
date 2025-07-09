// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/cluster/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datashare"
)

var datashareCrdGroup = "datashare.azure.upbound.io"

// Setup_datashare creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datashare(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: datashareCrdGroup, Kind: "Account"}:              account.Setup,
		schema.GroupKind{Group: datashareCrdGroup, Kind: "DataSetBlobStorage"}:   datasetblobstorage.Setup,
		schema.GroupKind{Group: datashareCrdGroup, Kind: "DataSetDataLakeGen2"}:  datasetdatalakegen2.Setup,
		schema.GroupKind{Group: datashareCrdGroup, Kind: "DataSetKustoCluster"}:  datasetkustocluster.Setup,
		schema.GroupKind{Group: datashareCrdGroup, Kind: "DataSetKustoDatabase"}: datasetkustodatabase.Setup,
		schema.GroupKind{Group: datashareCrdGroup, Kind: "DataShare"}:            datashare.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
