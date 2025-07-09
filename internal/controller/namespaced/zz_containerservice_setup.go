// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	kubernetescluster "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetescluster"
	kubernetesclusterextension "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesclusterextension"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesfleetmanager"
)

var containerserviceCrdGroup = "containerservice.azure.m.upbound.io"

// Setup_containerservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerservice(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: containerserviceCrdGroup, Kind: "KubernetesCluster"}:          kubernetescluster.Setup,
		schema.GroupKind{Group: containerserviceCrdGroup, Kind: "KubernetesClusterExtension"}: kubernetesclusterextension.Setup,
		schema.GroupKind{Group: containerserviceCrdGroup, Kind: "KubernetesClusterNodePool"}:  kubernetesclusternodepool.Setup,
		schema.GroupKind{Group: containerserviceCrdGroup, Kind: "KubernetesFleetManager"}:     kubernetesfleetmanager.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
