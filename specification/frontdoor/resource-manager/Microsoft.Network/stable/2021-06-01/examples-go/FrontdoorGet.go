package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorGet.json
func ExampleFrontDoorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFrontDoorsClient().Get(ctx, "rg1", "frontDoor1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FrontDoor = armfrontdoor.FrontDoor{
	// 	Name: to.Ptr("frontDoor1"),
	// 	Type: to.Ptr("Microsoft.Network/frontDoor"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armfrontdoor.Properties{
	// 		BackendPools: []*armfrontdoor.BackendPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
	// 				Name: to.Ptr("backendPool1"),
	// 				Properties: &armfrontdoor.BackendPoolProperties{
	// 					Backends: []*armfrontdoor.Backend{
	// 						{
	// 							Address: to.Ptr("w3.contoso.com"),
	// 							EnabledState: to.Ptr(armfrontdoor.BackendEnabledStateEnabled),
	// 							HTTPPort: to.Ptr[int32](80),
	// 							HTTPSPort: to.Ptr[int32](443),
	// 							Priority: to.Ptr[int32](2),
	// 							Weight: to.Ptr[int32](1),
	// 						},
	// 						{
	// 							Address: to.Ptr("contoso.com.website-us-west-2.othercloud.net"),
	// 							EnabledState: to.Ptr(armfrontdoor.BackendEnabledStateEnabled),
	// 							HTTPPort: to.Ptr[int32](80),
	// 							HTTPSPort: to.Ptr[int32](443),
	// 							Priority: to.Ptr[int32](1),
	// 							PrivateEndpointStatus: to.Ptr(armfrontdoor.PrivateEndpointStatusApproved),
	// 							PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
	// 							PrivateLinkLocation: to.Ptr("eastus"),
	// 							PrivateLinkResourceID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
	// 							Weight: to.Ptr[int32](2),
	// 						},
	// 						{
	// 							Address: to.Ptr("10.0.1.5"),
	// 							EnabledState: to.Ptr(armfrontdoor.BackendEnabledStateEnabled),
	// 							HTTPPort: to.Ptr[int32](80),
	// 							HTTPSPort: to.Ptr[int32](443),
	// 							Priority: to.Ptr[int32](1),
	// 							PrivateEndpointStatus: to.Ptr(armfrontdoor.PrivateEndpointStatusApproved),
	// 							PrivateLinkAlias: to.Ptr("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
	// 							PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
	// 							Weight: to.Ptr[int32](1),
	// 					}},
	// 					HealthProbeSettings: &armfrontdoor.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/healthProbeSettings/healthProbeSettings1"),
	// 					},
	// 					LoadBalancingSettings: &armfrontdoor.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/loadBalancingSettings/loadBalancingSettings1"),
	// 					},
	// 				},
	// 		}},
	// 		BackendPoolsSettings: &armfrontdoor.BackendPoolsSettings{
	// 			EnforceCertificateNameCheck: to.Ptr(armfrontdoor.EnforceCertificateNameCheckEnabledStateEnabled),
	// 			SendRecvTimeoutSeconds: to.Ptr[int32](60),
	// 		},
	// 		EnabledState: to.Ptr(armfrontdoor.FrontDoorEnabledStateEnabled),
	// 		FrontendEndpoints: []*armfrontdoor.FrontendEndpoint{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
	// 				Name: to.Ptr("frontendEndpoint1"),
	// 				Properties: &armfrontdoor.FrontendEndpointProperties{
	// 					HostName: to.Ptr("www.contoso.com"),
	// 					SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
	// 					SessionAffinityTTLSeconds: to.Ptr[int32](60),
	// 					WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/default"),
	// 				Name: to.Ptr("default"),
	// 				Properties: &armfrontdoor.FrontendEndpointProperties{
	// 					HostName: to.Ptr("frontDoor1.azurefd.net"),
	// 				},
	// 		}},
	// 		HealthProbeSettings: []*armfrontdoor.HealthProbeSettingsModel{
	// 			{
	// 				Name: to.Ptr("healthProbeSettings1"),
	// 				Properties: &armfrontdoor.HealthProbeSettingsProperties{
	// 					Path: to.Ptr("/"),
	// 					EnabledState: to.Ptr(armfrontdoor.HealthProbeEnabledEnabled),
	// 					HealthProbeMethod: to.Ptr(armfrontdoor.FrontDoorHealthProbeMethodHEAD),
	// 					IntervalInSeconds: to.Ptr[int32](120),
	// 					Protocol: to.Ptr(armfrontdoor.FrontDoorProtocolHTTP),
	// 				},
	// 		}},
	// 		LoadBalancingSettings: []*armfrontdoor.LoadBalancingSettingsModel{
	// 			{
	// 				Name: to.Ptr("loadBalancingSettings1"),
	// 				Properties: &armfrontdoor.LoadBalancingSettingsProperties{
	// 					SampleSize: to.Ptr[int32](4),
	// 					SuccessfulSamplesRequired: to.Ptr[int32](2),
	// 				},
	// 		}},
	// 		RoutingRules: []*armfrontdoor.RoutingRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/routingRules/routingRule1"),
	// 				Name: to.Ptr("routingRule1"),
	// 				Properties: &armfrontdoor.RoutingRuleProperties{
	// 					AcceptedProtocols: []*armfrontdoor.FrontDoorProtocol{
	// 						to.Ptr(armfrontdoor.FrontDoorProtocolHTTP)},
	// 						EnabledState: to.Ptr(armfrontdoor.RoutingRuleEnabledStateEnabled),
	// 						FrontendEndpoints: []*armfrontdoor.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/default"),
	// 						}},
	// 						PatternsToMatch: []*string{
	// 							to.Ptr("/*")},
	// 							RouteConfiguration: &armfrontdoor.ForwardingConfiguration{
	// 								ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration"),
	// 								BackendPool: &armfrontdoor.SubResource{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
	// 								},
	// 								CustomForwardingPath: to.Ptr(""),
	// 								ForwardingProtocol: to.Ptr(armfrontdoor.FrontDoorForwardingProtocolMatchRequest),
	// 							},
	// 							RulesEngine: &armfrontdoor.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
	// 							},
	// 							WebApplicationFirewallPolicyLink: &armfrontdoor.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
	// 							},
	// 						},
	// 				}},
	// 				Cname: to.Ptr("frontDoor1.azurefd.net"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				ResourceState: to.Ptr(armfrontdoor.FrontDoorResourceStateEnabled),
	// 				RulesEngines: []*armfrontdoor.RulesEngine{
	// 					{
	// 						Name: to.Ptr("rulesEngine1"),
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
	// 						Properties: &armfrontdoor.RulesEngineProperties{
	// 							Rules: []*armfrontdoor.RulesEngineRule{
	// 								{
	// 									Name: to.Ptr("Rule1"),
	// 									Action: &armfrontdoor.RulesEngineAction{
	// 										RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
	// 											ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration"),
	// 											CustomFragment: to.Ptr("fragment"),
	// 											CustomHost: to.Ptr("www.bing.com"),
	// 											CustomPath: to.Ptr("/api"),
	// 											CustomQueryString: to.Ptr("a=b"),
	// 											RedirectProtocol: to.Ptr(armfrontdoor.FrontDoorRedirectProtocolHTTPSOnly),
	// 											RedirectType: to.Ptr(armfrontdoor.FrontDoorRedirectTypeMoved),
	// 										},
	// 									},
	// 									MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
	// 										{
	// 											RulesEngineMatchValue: []*string{
	// 												to.Ptr("CH")},
	// 												RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRemoteAddr),
	// 												RulesEngineOperator: to.Ptr(armfrontdoor.RulesEngineOperatorGeoMatch),
	// 										}},
	// 										MatchProcessingBehavior: to.Ptr(armfrontdoor.MatchProcessingBehaviorStop),
	// 										Priority: to.Ptr[int32](1),
	// 									},
	// 									{
	// 										Name: to.Ptr("Rule2"),
	// 										Action: &armfrontdoor.RulesEngineAction{
	// 											ResponseHeaderActions: []*armfrontdoor.HeaderAction{
	// 												{
	// 													HeaderActionType: to.Ptr(armfrontdoor.HeaderActionTypeOverwrite),
	// 													HeaderName: to.Ptr("Cache-Control"),
	// 													Value: to.Ptr("public, max-age=31536000"),
	// 											}},
	// 										},
	// 										MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
	// 											{
	// 												RulesEngineMatchValue: []*string{
	// 													to.Ptr("jpg")},
	// 													RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariable("FilenameExtension")),
	// 													RulesEngineOperator: to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
	// 													Transforms: []*armfrontdoor.Transform{
	// 														to.Ptr(armfrontdoor.TransformLowercase)},
	// 												}},
	// 												Priority: to.Ptr[int32](2),
	// 										}},
	// 									},
	// 							}},
	// 						},
	// 					}
}
