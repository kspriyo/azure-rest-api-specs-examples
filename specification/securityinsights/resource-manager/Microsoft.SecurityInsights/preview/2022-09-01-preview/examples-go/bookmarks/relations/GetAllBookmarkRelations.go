package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/relations/GetAllBookmarkRelations.json
func ExampleBookmarkRelationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBookmarkRelationsClient().NewListPager("myRg", "myWorkspace", "2216d0e1-91e3-4902-89fd-d2df8c535096", &armsecurityinsights.BookmarkRelationsClientListOptions{Filter: nil,
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
		// 			Type: to.Ptr("Microsoft.SecurityInsights/bookmarks/relations"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/2216d0e1-91e3-4902-89fd-d2df8c535096/relations/4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
		// 			Etag: to.Ptr("190057d0-0000-0d00-0000-5c6f5adb0000"),
		// 			Properties: &armsecurityinsights.RelationProperties{
		// 				RelatedResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/afbd324f-6c48-459c-8710-8d1e1cd03812"),
		// 				RelatedResourceName: to.Ptr("afbd324f-6c48-459c-8710-8d1e1cd03812"),
		// 				RelatedResourceType: to.Ptr("Microsoft.SecurityInsights/incidents"),
		// 			},
		// 	}},
		// }
	}
}
