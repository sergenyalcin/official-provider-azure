// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	springcloudconnection "github.com/upbound/provider-azure/internal/controller/namespaced/servicelinker/springcloudconnection"
)

var servicelinkerCrdGroup = "servicelinker.azure.m.upbound.io"

// Setup_servicelinker creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicelinker(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: servicelinkerCrdGroup, Kind: "SpringCloudConnection"}: springcloudconnection.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, servicelinkerCrdGroup, o); err != nil {
		return err
	}
	return nil
}
