package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetRegistryValueEntityById.json
func ExampleEntitiesClient_Get_getARegistryValueEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "dc44bd11-b348-4d76-ad29-37bf7aa41356", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.RegistryValueEntity{
	// 		Name: to.Ptr("dc44bd11-b348-4d76-ad29-37bf7aa41356"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/dc44bd11-b348-4d76-ad29-37bf7aa41356"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindRegistryValue),
	// 		Properties: &armsecurityinsights.RegistryValueEntityProperties{
	// 			FriendlyName: to.Ptr("Data"),
	// 			KeyEntityID: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 			ValueData: to.Ptr("Data"),
	// 			ValueName: to.Ptr("Name"),
	// 			ValueType: to.Ptr(armsecurityinsights.RegistryValueKindString),
	// 		},
	// 	},
	// 	                        }
}
