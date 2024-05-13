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

func (mg *KubernetesCluster) ResolveReferences( // ResolveReferences of this KubernetesCluster.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.APIServerAccessProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIServerAccessProfile.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.APIServerAccessProfile.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.APIServerAccessProfile.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.APIServerAccessProfile.SubnetID")
		}
		mg.Spec.ForProvider.APIServerAccessProfile.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.APIServerAccessProfile.SubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.AciConnectorLinux != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AciConnectorLinux.SubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.AciConnectorLinux.SubnetNameRef,
				Selector:     mg.Spec.ForProvider.AciConnectorLinux.SubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AciConnectorLinux.SubnetName")
		}
		mg.Spec.ForProvider.AciConnectorLinux.SubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AciConnectorLinux.SubnetNameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.DefaultNodePool != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultNodePool.PodSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DefaultNodePool.PodSubnetIDRef,
				Selector:     mg.Spec.ForProvider.DefaultNodePool.PodSubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultNodePool.PodSubnetID")
		}
		mg.Spec.ForProvider.DefaultNodePool.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultNodePool.PodSubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.DefaultNodePool != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultNodePool.VnetSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DefaultNodePool.VnetSubnetIDRef,
				Selector:     mg.Spec.ForProvider.DefaultNodePool.VnetSubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultNodePool.VnetSubnetID")
		}
		mg.Spec.ForProvider.DefaultNodePool.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultNodePool.VnetSubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.IngressApplicationGateway != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IngressApplicationGateway.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.IngressApplicationGateway.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.IngressApplicationGateway.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IngressApplicationGateway.SubnetID")
		}
		mg.Spec.ForProvider.IngressApplicationGateway.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IngressApplicationGateway.SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "PrivateDNSZone", "PrivateDNSZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateDNSZoneID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrivateDNSZoneIDRef,
			Selector:     mg.Spec.ForProvider.PrivateDNSZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateDNSZoneID")
	}
	mg.Spec.ForProvider.PrivateDNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateDNSZoneIDRef = rsp.ResolvedReference
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

	if mg.Spec.InitProvider.APIServerAccessProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIServerAccessProfile.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.APIServerAccessProfile.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.APIServerAccessProfile.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.APIServerAccessProfile.SubnetID")
		}
		mg.Spec.InitProvider.APIServerAccessProfile.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.APIServerAccessProfile.SubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.AciConnectorLinux != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AciConnectorLinux.SubnetName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.AciConnectorLinux.SubnetNameRef,
				Selector:     mg.Spec.InitProvider.AciConnectorLinux.SubnetNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AciConnectorLinux.SubnetName")
		}
		mg.Spec.InitProvider.AciConnectorLinux.SubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AciConnectorLinux.SubnetNameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.DefaultNodePool != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultNodePool.PodSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DefaultNodePool.PodSubnetIDRef,
				Selector:     mg.Spec.InitProvider.DefaultNodePool.PodSubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultNodePool.PodSubnetID")
		}
		mg.Spec.InitProvider.DefaultNodePool.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultNodePool.PodSubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.DefaultNodePool != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultNodePool.VnetSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DefaultNodePool.VnetSubnetIDRef,
				Selector:     mg.Spec.InitProvider.DefaultNodePool.VnetSubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultNodePool.VnetSubnetID")
		}
		mg.Spec.InitProvider.DefaultNodePool.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultNodePool.VnetSubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.IngressApplicationGateway != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IngressApplicationGateway.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.IngressApplicationGateway.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.IngressApplicationGateway.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.IngressApplicationGateway.SubnetID")
		}
		mg.Spec.InitProvider.IngressApplicationGateway.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.IngressApplicationGateway.SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "PrivateDNSZone", "PrivateDNSZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateDNSZoneID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PrivateDNSZoneIDRef,
			Selector:     mg.Spec.InitProvider.PrivateDNSZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrivateDNSZoneID")
	}
	mg.Spec.InitProvider.PrivateDNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrivateDNSZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KubernetesClusterNodePool.
func (mg *KubernetesClusterNodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("containerservice.azure.upbound.io", "v1beta2", "KubernetesCluster", "KubernetesClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KubernetesClusterID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KubernetesClusterIDRef,
			Selector:     mg.Spec.ForProvider.KubernetesClusterIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KubernetesClusterID")
	}
	mg.Spec.ForProvider.KubernetesClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KubernetesClusterIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PodSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PodSubnetIDRef,
			Selector:     mg.Spec.ForProvider.PodSubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PodSubnetID")
	}
	mg.Spec.ForProvider.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PodSubnetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VnetSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VnetSubnetIDRef,
			Selector:     mg.Spec.ForProvider.VnetSubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VnetSubnetID")
	}
	mg.Spec.ForProvider.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VnetSubnetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PodSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PodSubnetIDRef,
			Selector:     mg.Spec.InitProvider.PodSubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PodSubnetID")
	}
	mg.Spec.InitProvider.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PodSubnetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VnetSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.VnetSubnetIDRef,
			Selector:     mg.Spec.InitProvider.VnetSubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VnetSubnetID")
	}
	mg.Spec.InitProvider.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VnetSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KubernetesFleetManager.
func (mg *KubernetesFleetManager) ResolveReferences(ctx context.Context, c client.Reader) error {
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
