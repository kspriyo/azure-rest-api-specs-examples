package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupUpdate.json
func ExampleSyncGroupsClient_BeginCreateOrUpdate_updateASyncGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginCreateOrUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", armsql.SyncGroup{
		Properties: &armsql.SyncGroupProperties{
			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
			HubDatabaseUserName:      to.Ptr("hubUser"),
			Interval:                 to.Ptr[int32](-1),
			SyncDatabaseID:           to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			UsePrivateLinkConnection: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncGroup = armsql.SyncGroup{
	// 	Name: to.Ptr("syncgroupcrud-3187"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 	Properties: &armsql.SyncGroupProperties{
	// 		ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 		HubDatabaseUserName: to.Ptr("hubUser"),
	// 		Interval: to.Ptr[int32](-1),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-3187"),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 		UsePrivateLinkConnection: to.Ptr(true),
	// 	},
	// }
}
