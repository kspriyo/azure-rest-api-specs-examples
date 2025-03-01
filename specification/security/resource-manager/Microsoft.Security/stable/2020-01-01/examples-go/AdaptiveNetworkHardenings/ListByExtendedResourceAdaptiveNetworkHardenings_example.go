package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/ListByExtendedResourceAdaptiveNetworkHardenings_example.json
func ExampleAdaptiveNetworkHardeningsClient_NewListByExtendedResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdaptiveNetworkHardeningsClient().NewListByExtendedResourcePager("rg1", "Microsoft.Compute", "virtualMachines", "vm1", nil)
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
		// page.AdaptiveNetworkHardeningsList = armsecurity.AdaptiveNetworkHardeningsList{
		// 	Value: []*armsecurity.AdaptiveNetworkHardening{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/adaptiveNetworkHardenings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Security/adaptiveNetworkHardenings/default"),
		// 			Properties: &armsecurity.AdaptiveNetworkHardeningProperties{
		// 				EffectiveNetworkSecurityGroups: []*armsecurity.EffectiveNetworkSecurityGroups{
		// 					{
		// 						NetworkInterface: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1"),
		// 						NetworkSecurityGroups: []*string{
		// 							to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nicNsg"),
		// 							to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/subnetNsg")},
		// 						},
		// 						{
		// 							NetworkInterface: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg1/providers/Microsoft.Network/networkInterfaces/nic2"),
		// 							NetworkSecurityGroups: []*string{
		// 								to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nicNsg2")},
		// 						}},
		// 						Rules: []*armsecurity.Rule{
		// 							{
		// 								Name: to.Ptr("rule1"),
		// 								DestinationPort: to.Ptr[int32](3389),
		// 								Direction: to.Ptr(armsecurity.DirectionInbound),
		// 								IPAddresses: []*string{
		// 									to.Ptr("100.10.1.1"),
		// 									to.Ptr("200.20.2.2"),
		// 									to.Ptr("81.199.3.0/24")},
		// 									Protocols: []*armsecurity.TransportProtocol{
		// 										to.Ptr(armsecurity.TransportProtocolTCP)},
		// 									},
		// 									{
		// 										Name: to.Ptr("rule2"),
		// 										DestinationPort: to.Ptr[int32](22),
		// 										Direction: to.Ptr(armsecurity.DirectionInbound),
		// 										IPAddresses: []*string{
		// 										},
		// 										Protocols: []*armsecurity.TransportProtocol{
		// 											to.Ptr(armsecurity.TransportProtocolTCP)},
		// 									}},
		// 									RulesCalculationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T00:00:00.000Z"); return t}()),
		// 								},
		// 						}},
		// 					}
	}
}
