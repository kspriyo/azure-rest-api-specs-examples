package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListDataWarehouseUserActivities.json
func ExampleDataWarehouseUserActivitiesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataWarehouseUserActivitiesClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb", nil)
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
		// page.DataWarehouseUserActivitiesListResult = armsql.DataWarehouseUserActivitiesListResult{
		// 	Value: []*armsql.DataWarehouseUserActivities{
		// 		{
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/dataWarehouseUserActivities"),
		// 			ID: to.Ptr("subscriptions/326affc3-21f4-4471-a545-e37430b70113/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/testsvr/databases/dwdb01/dataWarehouseUserActivities/current"),
		// 			Properties: &armsql.DataWarehouseUserActivitiesProperties{
		// 				ActiveQueriesCount: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
