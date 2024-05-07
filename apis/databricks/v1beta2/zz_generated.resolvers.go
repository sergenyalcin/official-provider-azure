// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AccessConnector) ResolveReferences( // ResolveReferences of this AccessConnector.
	ctx context.Context, c client.Reader) error {
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

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.CustomParameters != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomParameters.PrivateSubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CustomParameters.PrivateSubnetNameRef,
				Selector:     mg.Spec.ForProvider.CustomParameters.PrivateSubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomParameters.PrivateSubnetName")
		}
		mg.Spec.ForProvider.CustomParameters.PrivateSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomParameters.PrivateSubnetNameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.CustomParameters != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomParameters.PublicSubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CustomParameters.PublicSubnetNameRef,
				Selector:     mg.Spec.ForProvider.CustomParameters.PublicSubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomParameters.PublicSubnetName")
		}
		mg.Spec.ForProvider.CustomParameters.PublicSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomParameters.PublicSubnetNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ManagedResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ManagedResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedResourceGroupName")
	}
	mg.Spec.ForProvider.ManagedResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedResourceGroupNameRef = rsp.ResolvedReference
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

	if mg.Spec.InitProvider.CustomParameters != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CustomParameters.PrivateSubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.CustomParameters.PrivateSubnetNameRef,
				Selector:     mg.Spec.InitProvider.CustomParameters.PrivateSubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CustomParameters.PrivateSubnetName")
		}
		mg.Spec.InitProvider.CustomParameters.PrivateSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CustomParameters.PrivateSubnetNameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.CustomParameters != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CustomParameters.PublicSubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.CustomParameters.PublicSubnetNameRef,
				Selector:     mg.Spec.InitProvider.CustomParameters.PublicSubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CustomParameters.PublicSubnetName")
		}
		mg.Spec.InitProvider.CustomParameters.PublicSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CustomParameters.PublicSubnetNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ManagedResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ManagedResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ManagedResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ManagedResourceGroupName")
	}
	mg.Spec.InitProvider.ManagedResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ManagedResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
