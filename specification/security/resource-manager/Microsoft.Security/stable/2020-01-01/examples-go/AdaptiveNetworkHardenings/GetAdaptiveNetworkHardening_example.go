package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/GetAdaptiveNetworkHardening_example.json
func ExampleAdaptiveNetworkHardeningsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdaptiveNetworkHardeningsClient().Get(ctx, "rg1", "Microsoft.Compute", "virtualMachines", "vm1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdaptiveNetworkHardening = armsecurity.AdaptiveNetworkHardening{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/adaptiveNetworkHardenings"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Security/adaptiveNetworkHardenings/default"),
	// 	Properties: &armsecurity.AdaptiveNetworkHardeningProperties{
	// 		EffectiveNetworkSecurityGroups: []*armsecurity.EffectiveNetworkSecurityGroups{
	// 			{
	// 				NetworkInterface: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1"),
	// 				NetworkSecurityGroups: []*string{
	// 					to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nicNsg"),
	// 					to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/subnetNsg")},
	// 				},
	// 				{
	// 					NetworkInterface: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/rg2/providers/Microsoft.Network/networkInterfaces/nic2"),
	// 					NetworkSecurityGroups: []*string{
	// 						to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nicNsg")},
	// 				}},
	// 				Rules: []*armsecurity.Rule{
	// 					{
	// 						Name: to.Ptr("rule1"),
	// 						DestinationPort: to.Ptr[int32](3389),
	// 						Direction: to.Ptr(armsecurity.DirectionInbound),
	// 						IPAddresses: []*string{
	// 							to.Ptr("100.10.1.1"),
	// 							to.Ptr("200.20.2.2"),
	// 							to.Ptr("81.199.3.0/24")},
	// 							Protocols: []*armsecurity.TransportProtocol{
	// 								to.Ptr(armsecurity.TransportProtocolTCP)},
	// 							},
	// 							{
	// 								Name: to.Ptr("rule2"),
	// 								DestinationPort: to.Ptr[int32](22),
	// 								Direction: to.Ptr(armsecurity.DirectionInbound),
	// 								IPAddresses: []*string{
	// 								},
	// 								Protocols: []*armsecurity.TransportProtocol{
	// 									to.Ptr(armsecurity.TransportProtocolTCP)},
	// 							}},
	// 							RulesCalculationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T00:00:00.000Z"); return t}()),
	// 						},
	// 					}
}
