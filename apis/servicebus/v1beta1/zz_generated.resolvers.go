// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

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

func (mg *NamespaceAuthorizationRule) ResolveReferences( // ResolveReferences of this NamespaceAuthorizationRule.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceDisasterRecoveryConfig.
func (mg *NamespaceDisasterRecoveryConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "NamespaceAuthorizationRule", "NamespaceAuthorizationRuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AliasAuthorizationRuleID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AliasAuthorizationRuleIDRef,
			Selector:     mg.Spec.ForProvider.AliasAuthorizationRuleIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AliasAuthorizationRuleID")
	}
	mg.Spec.ForProvider.AliasAuthorizationRuleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AliasAuthorizationRuleIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PartnerNamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PartnerNamespaceIDRef,
			Selector:     mg.Spec.ForProvider.PartnerNamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PartnerNamespaceID")
	}
	mg.Spec.ForProvider.PartnerNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PartnerNamespaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrimaryNamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrimaryNamespaceIDRef,
			Selector:     mg.Spec.ForProvider.PrimaryNamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrimaryNamespaceID")
	}
	mg.Spec.ForProvider.PrimaryNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrimaryNamespaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "NamespaceAuthorizationRule", "NamespaceAuthorizationRuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AliasAuthorizationRuleID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.AliasAuthorizationRuleIDRef,
			Selector:     mg.Spec.InitProvider.AliasAuthorizationRuleIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AliasAuthorizationRuleID")
	}
	mg.Spec.InitProvider.AliasAuthorizationRuleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AliasAuthorizationRuleIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PartnerNamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PartnerNamespaceIDRef,
			Selector:     mg.Spec.InitProvider.PartnerNamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PartnerNamespaceID")
	}
	mg.Spec.InitProvider.PartnerNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PartnerNamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceNetworkRuleSet.
func (mg *NamespaceNetworkRuleSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRules[i3].SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.NetworkRules[i3].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.NetworkRules[i3].SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRules[i3].SubnetID")
		}
		mg.Spec.ForProvider.NetworkRules[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkRules[i3].SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.NamespaceIDRef,
			Selector:     mg.Spec.InitProvider.NamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NamespaceID")
	}
	mg.Spec.InitProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NamespaceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkRules[i3].SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.NetworkRules[i3].SubnetIDRef,
				Selector:     mg.Spec.InitProvider.NetworkRules[i3].SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkRules[i3].SubnetID")
		}
		mg.Spec.InitProvider.NetworkRules[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NetworkRules[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QueueAuthorizationRule.
func (mg *QueueAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.QueueIDRef,
			Selector:     mg.Spec.ForProvider.QueueIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueID")
	}
	mg.Spec.ForProvider.QueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceBusNamespace.
func (mg *ServiceBusNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkRuleSet); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDRef,
					Selector:     mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID")
			}
			mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
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

	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkRuleSet); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDRef,
					Selector:     mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID")
			}
			mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.NetworkRuleSet[i3].NetworkRules[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TopicIDRef,
			Selector:     mg.Spec.ForProvider.TopicIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicID")
	}
	mg.Spec.ForProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionRule.
func (mg *SubscriptionRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "Subscription", "SubscriptionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubscriptionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubscriptionIDRef,
			Selector:     mg.Spec.ForProvider.SubscriptionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubscriptionID")
	}
	mg.Spec.ForProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "ServiceBusNamespace", "ServiceBusNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicAuthorizationRule.
func (mg *TopicAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicebus.azure.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TopicIDRef,
			Selector:     mg.Spec.ForProvider.TopicIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicID")
	}
	mg.Spec.ForProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicIDRef = rsp.ResolvedReference

	return nil
}
