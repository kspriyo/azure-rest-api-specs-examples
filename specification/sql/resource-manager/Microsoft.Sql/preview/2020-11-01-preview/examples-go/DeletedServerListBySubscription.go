package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerListBySubscription.json
func ExampleDeletedServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedServersClient().NewListPager(nil)
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
		// page.DeletedServerListResult = armsql.DeletedServerListResult{
		// 	Value: []*armsql.DeletedServer{
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-1414"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-15T20:20:00.345Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-d-2424"),
		// 			Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-2424"),
		// 			Properties: &armsql.DeletedServerProperties{
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-13T10:10:00.678Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-2424.database.windows.net"),
		// 				OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-2424"),
		// 				Version: to.Ptr("12.0"),
		// 			},
		// 	}},
		// }
	}
}
