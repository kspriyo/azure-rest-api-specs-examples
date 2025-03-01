package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/GetBookmarks.json
func ExampleBookmarksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBookmarksClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.BookmarkList = armsecurityinsights.BookmarkList{
		// 	Value: []*armsecurityinsights.Bookmark{
		// 		{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/bookmarks"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.BookmarkProperties{
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t}()),
		// 				CreatedBy: &armsecurityinsights.UserInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 				},
		// 				DisplayName: to.Ptr("My bookmark"),
		// 				EntityMappings: []*armsecurityinsights.BookmarkEntityMappings{
		// 					{
		// 						EntityType: to.Ptr("Account"),
		// 						FieldMappings: []*armsecurityinsights.EntityFieldMapping{
		// 							{
		// 								Identifier: to.Ptr("Fullname"),
		// 								Value: to.Ptr("johndoe@microsoft.com"),
		// 						}},
		// 				}},
		// 				IncidentInfo: &armsecurityinsights.IncidentInfo{
		// 					IncidentID: to.Ptr("DDA55F97-170B-40B9-B8ED-CBFD05481E7D"),
		// 					RelationName: to.Ptr("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0018"),
		// 					Severity: to.Ptr(armsecurityinsights.IncidentSeverityLow),
		// 					Title: to.Ptr("New case 1"),
		// 				},
		// 				Labels: []*string{
		// 					to.Ptr("Tag1"),
		// 					to.Ptr("Tag2")},
		// 					Notes: to.Ptr("Found a suspicious activity"),
		// 					Query: to.Ptr("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)"),
		// 					QueryResult: to.Ptr("Security Event query result"),
		// 					Tactics: []*armsecurityinsights.AttackTactic{
		// 						to.Ptr(armsecurityinsights.AttackTacticExecution)},
		// 						Techniques: []*string{
		// 							to.Ptr("T1609")},
		// 							Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t}()),
		// 							UpdatedBy: &armsecurityinsights.UserInfo{
		// 								Name: to.Ptr("john doe"),
		// 								Email: to.Ptr("john@contoso.com"),
		// 								ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
