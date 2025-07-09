// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	computecluster "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/synapsespark"
	workspace "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/workspace"
)

var machinelearningservicesCrdGroup = "machinelearningservices.azure.upbound.io"

// Setup_machinelearningservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_machinelearningservices(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: machinelearningservicesCrdGroup, Kind: "ComputeCluster"}:  computecluster.Setup,
		schema.GroupKind{Group: machinelearningservicesCrdGroup, Kind: "ComputeInstance"}: computeinstance.Setup,
		schema.GroupKind{Group: machinelearningservicesCrdGroup, Kind: "SynapseSpark"}:    synapsespark.Setup,
		schema.GroupKind{Group: machinelearningservicesCrdGroup, Kind: "Workspace"}:       workspace.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
