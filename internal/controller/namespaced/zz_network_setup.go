// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	applicationgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/applicationsecuritygroup"
	bastionhost "github.com/upbound/provider-azure/internal/controller/namespaced/network/bastionhost"
	connectionmonitor "github.com/upbound/provider-azure/internal/controller/namespaced/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/internal/controller/namespaced/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/ipgroup"
	loadbalancer "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/localnetworkgateway"
	manager "github.com/upbound/provider-azure/internal/controller/namespaced/network/manager"
	managermanagementgroupconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/managernetworkgroup"
	managerstaticmember "github.com/upbound/provider-azure/internal/controller/namespaced/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/managersubscriptionconnection"
	natgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/internal/controller/namespaced/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsptrrecord"
	privatednsresolver "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolver"
	privatednsresolverdnsforwardingruleset "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverdnsforwardingruleset"
	privatednsresolverforwardingrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverforwardingrule"
	privatednsresolverinboundendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverinboundendpoint"
	privatednsresolveroutboundendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolveroutboundendpoint"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatelinkservice"
	profile "github.com/upbound/provider-azure/internal/controller/namespaced/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/namespaced/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/namespaced/network/publicipprefix"
	route "github.com/upbound/provider-azure/internal/controller/namespaced/network/route"
	routefilter "github.com/upbound/provider-azure/internal/controller/namespaced/network/routefilter"
	routemap "github.com/upbound/provider-azure/internal/controller/namespaced/network/routemap"
	routeserver "github.com/upbound/provider-azure/internal/controller/namespaced/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/internal/controller/namespaced/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubroutetableroute"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubsecuritypartnerprovider"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnserverconfigurationpolicygroup"
	vpnsite "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnsite"
	watcher "github.com/upbound/provider-azure/internal/controller/namespaced/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/namespaced/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/webapplicationfirewallpolicy"
)

var networkCrdGroup = "network.azure.m.upbound.io"

// Setup_network creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_network(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: networkCrdGroup, Kind: "ApplicationGateway"}:                                  applicationgateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ApplicationSecurityGroup"}:                            applicationsecuritygroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "BastionHost"}:                                         bastionhost.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ConnectionMonitor"}:                                   connectionmonitor.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DDoSProtectionPlan"}:                                  ddosprotectionplan.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSAAAARecord"}:                                       dnsaaaarecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSARecord"}:                                          dnsarecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSCAARecord"}:                                        dnscaarecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSCNAMERecord"}:                                      dnscnamerecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSMXRecord"}:                                         dnsmxrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSNSRecord"}:                                         dnsnsrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSPTRRecord"}:                                        dnsptrrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSSRVRecord"}:                                        dnssrvrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSTXTRecord"}:                                        dnstxtrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "DNSZone"}:                                             dnszone.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteCircuit"}:                                 expressroutecircuit.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteCircuitAuthorization"}:                    expressroutecircuitauthorization.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteCircuitConnection"}:                       expressroutecircuitconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteCircuitPeering"}:                          expressroutecircuitpeering.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteConnection"}:                              expressrouteconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRouteGateway"}:                                 expressroutegateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ExpressRoutePort"}:                                    expressrouteport.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Firewall"}:                                            firewall.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FirewallApplicationRuleCollection"}:                   firewallapplicationrulecollection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FirewallNATRuleCollection"}:                           firewallnatrulecollection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FirewallNetworkRuleCollection"}:                       firewallnetworkrulecollection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FirewallPolicy"}:                                      firewallpolicy.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FirewallPolicyRuleCollectionGroup"}:                   firewallpolicyrulecollectiongroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FrontDoor"}:                                           frontdoor.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FrontdoorCustomHTTPSConfiguration"}:                   frontdoorcustomhttpsconfiguration.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FrontdoorFirewallPolicy"}:                             frontdoorfirewallpolicy.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "FrontdoorRulesEngine"}:                                frontdoorrulesengine.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "IPGroup"}:                                             ipgroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancer"}:                                        loadbalancer.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerBackendAddressPool"}:                      loadbalancerbackendaddresspool.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerBackendAddressPoolAddress"}:               loadbalancerbackendaddresspooladdress.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerNatPool"}:                                 loadbalancernatpool.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerNatRule"}:                                 loadbalancernatrule.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerOutboundRule"}:                            loadbalanceroutboundrule.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerProbe"}:                                   loadbalancerprobe.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LoadBalancerRule"}:                                    loadbalancerrule.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "LocalNetworkGateway"}:                                 localnetworkgateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Manager"}:                                             manager.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ManagerManagementGroupConnection"}:                    managermanagementgroupconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ManagerNetworkGroup"}:                                 managernetworkgroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ManagerStaticMember"}:                                 managerstaticmember.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "ManagerSubscriptionConnection"}:                       managersubscriptionconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NATGateway"}:                                          natgateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NATGatewayPublicIPAssociation"}:                       natgatewaypublicipassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NATGatewayPublicIPPrefixAssociation"}:                 natgatewaypublicipprefixassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NetworkInterface"}:                                    networkinterface.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NetworkInterfaceApplicationSecurityGroupAssociation"}: networkinterfaceapplicationsecuritygroupassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NetworkInterfaceBackendAddressPoolAssociation"}:       networkinterfacebackendaddresspoolassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NetworkInterfaceNatRuleAssociation"}:                  networkinterfacenatruleassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "NetworkInterfaceSecurityGroupAssociation"}:            networkinterfacesecuritygroupassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PacketCapture"}:                                       packetcapture.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PointToSiteVPNGateway"}:                               pointtositevpngateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSAAAARecord"}:                                privatednsaaaarecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSARecord"}:                                   privatednsarecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSCNAMERecord"}:                               privatednscnamerecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSMXRecord"}:                                  privatednsmxrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSPTRRecord"}:                                 privatednsptrrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSResolver"}:                                  privatednsresolver.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSResolverDNSForwardingRuleset"}:              privatednsresolverdnsforwardingruleset.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSResolverForwardingRule"}:                    privatednsresolverforwardingrule.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSResolverInboundEndpoint"}:                   privatednsresolverinboundendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSResolverOutboundEndpoint"}:                  privatednsresolveroutboundendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSSRVRecord"}:                                 privatednssrvrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSTXTRecord"}:                                 privatednstxtrecord.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSZone"}:                                      privatednszone.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateDNSZoneVirtualNetworkLink"}:                    privatednszonevirtualnetworklink.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateEndpoint"}:                                     privateendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateEndpointApplicationSecurityGroupAssociation"}:  privateendpointapplicationsecuritygroupassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PrivateLinkService"}:                                  privatelinkservice.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Profile"}:                                             profile.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PublicIP"}:                                            publicip.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "PublicIPPrefix"}:                                      publicipprefix.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Route"}:                                               route.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "RouteFilter"}:                                         routefilter.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "RouteMap"}:                                            routemap.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "RouteServer"}:                                         routeserver.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "RouteServerBGPConnection"}:                            routeserverbgpconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "RouteTable"}:                                          routetable.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SecurityGroup"}:                                       securitygroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SecurityRule"}:                                        securityrule.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Subnet"}:                                              subnet.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SubnetNATGatewayAssociation"}:                         subnetnatgatewayassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SubnetNetworkSecurityGroupAssociation"}:               subnetnetworksecuritygroupassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SubnetRouteTableAssociation"}:                         subnetroutetableassociation.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "SubnetServiceEndpointStoragePolicy"}:                  subnetserviceendpointstoragepolicy.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "TrafficManagerAzureEndpoint"}:                         trafficmanagerazureendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "TrafficManagerExternalEndpoint"}:                      trafficmanagerexternalendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "TrafficManagerNestedEndpoint"}:                        trafficmanagernestedendpoint.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "TrafficManagerProfile"}:                               trafficmanagerprofile.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VPNGateway"}:                                          vpngateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VPNGatewayConnection"}:                                vpngatewayconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VPNServerConfiguration"}:                              vpnserverconfiguration.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VPNServerConfigurationPolicyGroup"}:                   vpnserverconfigurationpolicygroup.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VPNSite"}:                                             vpnsite.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHub"}:                                          virtualhub.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHubConnection"}:                                virtualhubconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHubIP"}:                                        virtualhubip.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHubRouteTable"}:                                virtualhubroutetable.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHubRouteTableRoute"}:                           virtualhubroutetableroute.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualHubSecurityPartnerProvider"}:                   virtualhubsecuritypartnerprovider.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualNetwork"}:                                      virtualnetwork.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualNetworkGateway"}:                               virtualnetworkgateway.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualNetworkGatewayConnection"}:                     virtualnetworkgatewayconnection.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualNetworkPeering"}:                               virtualnetworkpeering.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "VirtualWAN"}:                                          virtualwan.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "Watcher"}:                                             watcher.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "WatcherFlowLog"}:                                      watcherflowlog.Setup,
		schema.GroupKind{Group: networkCrdGroup, Kind: "WebApplicationFirewallPolicy"}:                        webapplicationfirewallpolicy.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, networkCrdGroup, o); err != nil {
		return err
	}
	return nil
}
