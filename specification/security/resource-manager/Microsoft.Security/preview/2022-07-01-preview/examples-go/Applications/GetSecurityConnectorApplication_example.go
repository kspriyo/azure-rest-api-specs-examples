package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/GetSecurityConnectorApplication_example.json
func ExampleConnectorApplicationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorApplicationClient().Get(ctx, "gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armsecurity.Application{
	// 	Name: to.Ptr("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
	// 	Type: to.Ptr("Microsoft.Security/applications"),
	// 	ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector/providers/Microsoft.Security/applications/ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
	// 	Properties: &armsecurity.ApplicationProperties{
	// 		Description: to.Ptr("An application on critical GCP recommendations"),
	// 		ConditionSets: []any{
	// 			map[string]any{
	// 				"conditions":[]any{
	// 					map[string]any{
	// 						"operator": "contains",
	// 						"property": "$.Id",
	// 						"value": "-bil-",
	// 					},
	// 				},
	// 		}},
	// 		DisplayName: to.Ptr("GCP Admin's application"),
	// 		SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
	// 	},
	// }
}
