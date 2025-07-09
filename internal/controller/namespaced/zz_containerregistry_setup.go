// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	agentpool "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/token"
	tokenpassword "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/tokenpassword"
	webhook "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/webhook"
)

var containerregistryCrdGroup = "containerregistry.azure.m.upbound.io"

// Setup_containerregistry creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerregistry(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "AgentPool"}:                  agentpool.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "ContainerConnectedRegistry"}: containerconnectedregistry.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "Registry"}:                   registry.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "ScopeMap"}:                   scopemap.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "Token"}:                      token.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "TokenPassword"}:              tokenpassword.Setup,
		schema.GroupKind{Group: containerregistryCrdGroup, Kind: "Webhook"}:                    webhook.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
