// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	rediscache "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscache"
	rediscacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicy"
	rediscacheaccesspolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicyassignment"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cluster/cache/redislinkedserver"
)

var cacheCrdGroup = "cache.azure.upbound.io"

// Setup_cache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cache(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisCache"}:                       rediscache.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisCacheAccessPolicy"}:           rediscacheaccesspolicy.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisCacheAccessPolicyAssignment"}: rediscacheaccesspolicyassignment.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisEnterpriseCluster"}:           redisenterprisecluster.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisEnterpriseDatabase"}:          redisenterprisedatabase.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisFirewallRule"}:                redisfirewallrule.Setup,
		schema.GroupKind{Group: cacheCrdGroup, Kind: "RedisLinkedServer"}:                redislinkedserver.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, cacheCrdGroup, o); err != nil {
		return err
	}
	return nil
}
