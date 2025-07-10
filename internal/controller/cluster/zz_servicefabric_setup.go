// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	cluster "github.com/upbound/provider-azure/internal/controller/cluster/servicefabric/cluster"
	managedcluster "github.com/upbound/provider-azure/internal/controller/cluster/servicefabric/managedcluster"
)

var servicefabricCrdGroup = "servicefabric.azure.upbound.io"

// Setup_servicefabric creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicefabric(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: servicefabricCrdGroup, Kind: "Cluster"}:        cluster.Setup,
		schema.GroupKind{Group: servicefabricCrdGroup, Kind: "ManagedCluster"}: managedcluster.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, servicefabricCrdGroup, o); err != nil {
		return err
	}
	return nil
}
