package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/PutResourcePricingByNameVirtualMachines_example.json
func ExamplePricingsClient_Update_updatePricingOnResourceExampleForVirtualMachinesPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().Update(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1", "virtualMachines", armsecurity.Pricing{
		Properties: &armsecurity.PricingProperties{
			PricingTier: to.Ptr(armsecurity.PricingTierStandard),
			SubPlan:     to.Ptr("P1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pricing = armsecurity.Pricing{
	// 	Name: to.Ptr("virtualMachines"),
	// 	Type: to.Ptr("Microsoft.Security/pricings"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1/providers/Microsoft.Security/pricings/virtualMachines"),
	// 	Properties: &armsecurity.PricingProperties{
	// 		EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 		FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 		Inherited: to.Ptr(armsecurity.InheritedFalse),
	// 		PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 		SubPlan: to.Ptr("P1"),
	// 		Extensions: []*armsecurity.Extension{
	// 			{
	// 				Name: to.Ptr("MdeDesignatedSubscription"),
	// 				IsEnabled: to.Ptr(armsecurity.IsEnabledFalse),
	// 			},
	// 			{
	// 				Name: to.Ptr("AgentlessVmScanning"),
	// 				AdditionalExtensionProperties: map[string]any{
	// 					"ExclusionTags": "[{\"Key\":\"TestKey1\",\"Value\":\"TestValue1\"},{\"Key\":\"TestKey2\",\"Value\":\"TestValue2\"}]",
	// 				},
	// 				IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 		}},
	// 	},
	// }
}
