// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	globalvmshutdownschedule "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/lab"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/linuxvirtualmachine"
	policy "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/policy"
	schedule "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/schedule"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/virtualnetwork"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/windowsvirtualmachine"
)

var devtestlabCrdGroup = "devtestlab.azure.m.upbound.io"

// Setup_devtestlab creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devtestlab(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "GlobalVMShutdownSchedule"}: globalvmshutdownschedule.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "Lab"}:                      lab.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "LinuxVirtualMachine"}:      linuxvirtualmachine.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "Policy"}:                   policy.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "Schedule"}:                 schedule.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "VirtualNetwork"}:           virtualnetwork.Setup,
		schema.GroupKind{Group: devtestlabCrdGroup, Kind: "WindowsVirtualMachine"}:    windowsvirtualmachine.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
