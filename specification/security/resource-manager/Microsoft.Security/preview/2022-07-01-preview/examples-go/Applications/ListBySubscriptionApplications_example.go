package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/ListBySubscriptionApplications_example.json
func ExampleApplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListPager(nil)
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
		// page.ApplicationsList = armsecurity.ApplicationsList{
		// 	Value: []*armsecurity.Application{
		// 		{
		// 			Name: to.Ptr("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
		// 			Type: to.Ptr("Microsoft.Security/applications"),
		// 			ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/applications/ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
		// 			Properties: &armsecurity.ApplicationProperties{
		// 				Description: to.Ptr("An application on critical recommendations"),
		// 				ConditionSets: []any{
		// 					map[string]any{
		// 						"conditions":[]any{
		// 							map[string]any{
		// 								"operator": "contains",
		// 								"property": "$.Id",
		// 								"value": "-bil-",
		// 							},
		// 						},
		// 				}},
		// 				DisplayName: to.Ptr("Admin's application"),
		// 				SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4106f43c-6d82-4fc8-a92c-dcfe50799d1d"),
		// 			Type: to.Ptr("Microsoft.Security/applications"),
		// 			ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/applications/4106f43c-6d82-4fc8-a92c-dcfe50799d1d"),
		// 			Properties: &armsecurity.ApplicationProperties{
		// 				Description: to.Ptr("An application on critical recommendations"),
		// 				ConditionSets: []any{
		// 					map[string]any{
		// 						"conditions":[]any{
		// 							map[string]any{
		// 								"operator": "contains",
		// 								"property": "$.Id",
		// 								"value": "-prod-",
		// 							},
		// 						},
		// 				}},
		// 				DisplayName: to.Ptr("Admin's application"),
		// 				SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
		// 			},
		// 	}},
		// }
	}
}
