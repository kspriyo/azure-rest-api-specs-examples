package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/relations/GetAllIncidentRelations.json
func ExampleIncidentRelationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIncidentRelationsClient().NewListPager("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", &armsecurityinsights.IncidentRelationsClientListOptions{Filter: nil,
		Orderby:   nil,
		Top:       nil,
		SkipToken: nil,
	})
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
		// page.RelationList = armsecurityinsights.RelationList{
		// 	Value: []*armsecurityinsights.Relation{
		// 		{
		// 			Name: to.Ptr("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/incidents/relations"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/afbd324f-6c48-459c-8710-8d1e1cd03812/relations/4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
		// 			Etag: to.Ptr("190057d0-0000-0d00-0000-5c6f5adb0000"),
		// 			Properties: &armsecurityinsights.RelationProperties{
		// 				RelatedResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 				RelatedResourceName: to.Ptr("2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 				RelatedResourceType: to.Ptr("Microsoft.SecurityInsights/bookmarks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("9673a17d-8bc7-4ca6-88ee-38a4f3efc032"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/incidents/relations"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/afbd324f-6c48-459c-8710-8d1e1cd03812/relations/9673a17d-8bc7-4ca6-88ee-38a4f3efc032"),
		// 			Etag: to.Ptr("6f714025-dd7c-46aa-b5d0-b9857488d060"),
		// 			Properties: &armsecurityinsights.RelationProperties{
		// 				RelatedResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/1dd267cd-8a1f-4f6f-b92c-da43ac8819af"),
		// 				RelatedResourceKind: to.Ptr("SecurityAlert"),
		// 				RelatedResourceName: to.Ptr("1dd267cd-8a1f-4f6f-b92c-da43ac8819af"),
		// 				RelatedResourceType: to.Ptr("Microsoft.SecurityInsights/entities"),
		// 			},
		// 	}},
		// }
	}
}
