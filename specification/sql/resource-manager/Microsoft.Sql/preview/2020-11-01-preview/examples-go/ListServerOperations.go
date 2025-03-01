package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListServerOperations.json
func ExampleServerOperationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerOperationsClient().NewListByServerPager("sqlcrudtest-7398", "sqlcrudtest-4645", nil)
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
		// page.ServerOperationListResult = armsql.ServerOperationListResult{
		// 	Value: []*armsql.ServerOperation{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/operations/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armsql.ServerOperationProperties{
		// 				Operation: to.Ptr("MakeAllLogicalDatabasesAccessible"),
		// 				OperationFriendlyName: to.Ptr("MAKE ALL DBS ACCESSIBLE"),
		// 				PercentComplete: to.Ptr[int32](0),
		// 				ServerName: to.Ptr("sqlcrudtest-4645"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T09:10:08.100Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateInProgress),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/operations/22222222-2222-2222-2222-222222222222"),
		// 			Properties: &armsql.ServerOperationProperties{
		// 				Operation: to.Ptr("MakeAllLogicalDatabasesAccessible"),
		// 				OperationFriendlyName: to.Ptr("MAKE ALL DBS ACCESSIBLE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("sqlcrudtest-4645"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T09:00:08.100Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
