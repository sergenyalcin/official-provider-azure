// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/namespaced/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterservervulnerabilityassessment"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterworkspace"
	storagedefender "github.com/upbound/provider-azure/internal/controller/namespaced/security/storagedefender"
)

var securityCrdGroup = "security.azure.m.upbound.io"

// Setup_security creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_security(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: securityCrdGroup, Kind: "AdvancedThreatProtection"}:                                  advancedthreatprotection.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "IOTSecurityDeviceGroup"}:                                    iotsecuritydevicegroup.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "IOTSecuritySolution"}:                                       iotsecuritysolution.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterAssessment"}:                                  securitycenterassessment.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterAssessmentPolicy"}:                            securitycenterassessmentpolicy.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterAutoProvisioning"}:                            securitycenterautoprovisioning.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterContact"}:                                     securitycentercontact.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterServerVulnerabilityAssessment"}:               securitycenterservervulnerabilityassessment.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterServerVulnerabilityAssessmentVirtualMachine"}: securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterSetting"}:                                     securitycentersetting.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterSubscriptionPricing"}:                         securitycentersubscriptionpricing.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "SecurityCenterWorkspace"}:                                   securitycenterworkspace.Setup,
		schema.GroupKind{Group: securityCrdGroup, Kind: "StorageDefender"}:                                           storagedefender.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, securityCrdGroup, o); err != nil {
		return err
	}
	return nil
}
