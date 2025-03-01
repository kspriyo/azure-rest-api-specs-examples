package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineList.json
func ExampleRulesEnginesClient_NewListByFrontDoorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRulesEnginesClient().NewListByFrontDoorPager("rg1", "frontDoor1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RulesEngineListResult = armfrontdoor.RulesEngineListResult{
		// 	Value: []*armfrontdoor.RulesEngine{
		// 		{
		// 			Name: to.Ptr("rulesEngine1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
		// 			Properties: &armfrontdoor.RulesEngineProperties{
		// 				Rules: []*armfrontdoor.RulesEngineRule{
		// 					{
		// 						Name: to.Ptr("Rule1"),
		// 						Action: &armfrontdoor.RulesEngineAction{
		// 							RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
		// 								ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration"),
		// 								CustomFragment: to.Ptr("fragment"),
		// 								CustomHost: to.Ptr("www.bing.com"),
		// 								CustomPath: to.Ptr("/api"),
		// 								CustomQueryString: to.Ptr("a=b"),
		// 								RedirectProtocol: to.Ptr(armfrontdoor.FrontDoorRedirectProtocolHTTPSOnly),
		// 								RedirectType: to.Ptr(armfrontdoor.FrontDoorRedirectTypeMoved),
		// 							},
		// 						},
		// 						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
		// 							{
		// 								RulesEngineMatchValue: []*string{
		// 									to.Ptr("CH")},
		// 									RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRemoteAddr),
		// 									RulesEngineOperator: to.Ptr(armfrontdoor.RulesEngineOperatorGeoMatch),
		// 							}},
		// 							MatchProcessingBehavior: to.Ptr(armfrontdoor.MatchProcessingBehaviorStop),
		// 							Priority: to.Ptr[int32](1),
		// 						},
		// 						{
		// 							Name: to.Ptr("Rule2"),
		// 							Action: &armfrontdoor.RulesEngineAction{
		// 								ResponseHeaderActions: []*armfrontdoor.HeaderAction{
		// 									{
		// 										HeaderActionType: to.Ptr(armfrontdoor.HeaderActionTypeOverwrite),
		// 										HeaderName: to.Ptr("Cache-Control"),
		// 										Value: to.Ptr("public, max-age=31536000"),
		// 								}},
		// 							},
		// 							MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
		// 								{
		// 									RulesEngineMatchValue: []*string{
		// 										to.Ptr("jpg")},
		// 										RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestFilenameExtension),
		// 										RulesEngineOperator: to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
		// 										Transforms: []*armfrontdoor.Transform{
		// 											to.Ptr(armfrontdoor.TransformLowercase)},
		// 									}},
		// 									Priority: to.Ptr[int32](2),
		// 								},
		// 								{
		// 									Name: to.Ptr("Rule3"),
		// 									Action: &armfrontdoor.RulesEngineAction{
		// 										RouteConfigurationOverride: &armfrontdoor.ForwardingConfiguration{
		// 											ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration"),
		// 											BackendPool: &armfrontdoor.SubResource{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
		// 											},
		// 											CacheConfiguration: &armfrontdoor.CacheConfiguration{
		// 												CacheDuration: to.Ptr("P1DT12H20M30S"),
		// 												DynamicCompression: to.Ptr(armfrontdoor.DynamicCompressionEnabledDisabled),
		// 												QueryParameterStripDirective: to.Ptr(armfrontdoor.FrontDoorQueryStripOnly),
		// 												QueryParameters: to.Ptr("a=b,p=q"),
		// 											},
		// 											ForwardingProtocol: to.Ptr(armfrontdoor.FrontDoorForwardingProtocolHTTPSOnly),
		// 										},
		// 									},
		// 									MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
		// 										{
		// 											NegateCondition: to.Ptr(false),
		// 											RulesEngineMatchValue: []*string{
		// 												to.Ptr("allowoverride")},
		// 												RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestHeader),
		// 												RulesEngineOperator: to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
		// 												Selector: to.Ptr("Rules-Engine-Route-Forward"),
		// 												Transforms: []*armfrontdoor.Transform{
		// 													to.Ptr(armfrontdoor.TransformLowercase)},
		// 											}},
		// 											Priority: to.Ptr[int32](3),
		// 									}},
		// 								},
		// 						}},
		// 					}
	}
}
