// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	domain "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/systemtopic"
	topic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/topic"
)

var eventgridCrdGroup = "eventgrid.azure.m.upbound.io"

// Setup_eventgrid creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventgrid(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: eventgridCrdGroup, Kind: "Domain"}:            domain.Setup,
		schema.GroupKind{Group: eventgridCrdGroup, Kind: "DomainTopic"}:       domaintopic.Setup,
		schema.GroupKind{Group: eventgridCrdGroup, Kind: "EventSubscription"}: eventsubscription.Setup,
		schema.GroupKind{Group: eventgridCrdGroup, Kind: "SystemTopic"}:       systemtopic.Setup,
		schema.GroupKind{Group: eventgridCrdGroup, Kind: "Topic"}:             topic.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, eventgridCrdGroup, o); err != nil {
		return err
	}
	return nil
}
