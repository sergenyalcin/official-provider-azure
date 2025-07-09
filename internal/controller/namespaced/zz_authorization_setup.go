// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	managementgrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/trustedaccessrolebinding"
)

var authorizationCrdGroup = "authorization.azure.m.upbound.io"

// Setup_authorization creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_authorization(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ManagementGroupPolicyAssignment"}: managementgrouppolicyassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ManagementGroupPolicyExemption"}:  managementgrouppolicyexemption.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ManagementLock"}:                  managementlock.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "PimActiveRoleAssignment"}:         pimactiveroleassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "PimEligibleRoleAssignment"}:       pimeligibleroleassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "PolicyDefinition"}:                policydefinition.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "PolicySetDefinition"}:             policysetdefinition.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ResourceGroupPolicyAssignment"}:   resourcegrouppolicyassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ResourceGroupPolicyExemption"}:    resourcegrouppolicyexemption.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ResourcePolicyAssignment"}:        resourcepolicyassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "ResourcePolicyExemption"}:         resourcepolicyexemption.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "RoleAssignment"}:                  roleassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "RoleDefinition"}:                  roledefinition.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "SubscriptionPolicyAssignment"}:    subscriptionpolicyassignment.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "SubscriptionPolicyExemption"}:     subscriptionpolicyexemption.Setup,
		schema.GroupKind{Group: authorizationCrdGroup, Kind: "TrustedAccessRoleBinding"}:        trustedaccessrolebinding.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
