// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AttachedDatabaseConfiguration) ResolveReferences( // ResolveReferences of this AttachedDatabaseConfiguration.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ClusterResourceIDRef,
			Selector:     mg.Spec.ForProvider.ClusterResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterResourceID")
	}
	mg.Spec.ForProvider.ClusterResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DatabaseNameRef,
			Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ClusterResourceIDRef,
			Selector:     mg.Spec.InitProvider.ClusterResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterResourceID")
	}
	mg.Spec.InitProvider.ClusterResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatabaseName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DatabaseNameRef,
			Selector:     mg.Spec.InitProvider.DatabaseNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatabaseName")
	}
	mg.Spec.InitProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatabaseNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.VirtualNetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID")
		}
		mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.VirtualNetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID")
		}
		mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}
