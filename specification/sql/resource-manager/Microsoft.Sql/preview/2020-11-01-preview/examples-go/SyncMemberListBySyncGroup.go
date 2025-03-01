package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberListBySyncGroup.json
func ExampleSyncMembersClient_NewListBySyncGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncMembersClient().NewListBySyncGroupPager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
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
		// page.SyncMemberListResult = armsql.SyncMemberListResult{
		// 	Value: []*armsql.SyncMember{
		// 		{
		// 			Name: to.Ptr("syncmembercrud-4879"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups/syncMembers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187/syncMembers/syncmembercrud-4879"),
		// 			Properties: &armsql.SyncMemberProperties{
		// 				DatabaseName: to.Ptr("syncgroupcrud-7421"),
		// 				DatabaseType: to.Ptr(armsql.SyncMemberDbTypeAzureSQLDatabase),
		// 				PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncmembercrud-4879"),
		// 				ServerName: to.Ptr("syncgroupcrud-3379.database.windows.net"),
		// 				SyncDirection: to.Ptr(armsql.SyncDirectionBidirectional),
		// 				SyncMemberAzureDatabaseResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
		// 				SyncState: to.Ptr(armsql.SyncMemberStateUnProvisioned),
		// 				UsePrivateLinkConnection: to.Ptr(true),
		// 				UserName: to.Ptr("myUser"),
		// 			},
		// 	}},
		// }
	}
}
