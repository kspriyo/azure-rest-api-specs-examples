package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedDatabaseListByManagedInstance.json
func ExampleManagedDatabasesClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabasesClient().NewListByInstancePager("Test1", "managedInstance", nil)
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
		// page.ManagedDatabaseListResult = armsql.ManagedDatabaseListResult{
		// 	Value: []*armsql.ManagedDatabase{
		// 		{
		// 			Name: to.Ptr("testdb1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/databases/testdb1"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Properties: &armsql.ManagedDatabaseProperties{
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-04T15:00:17.730Z"); return t}()),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				IsLedgerOn: to.Ptr(false),
		// 				Status: to.Ptr(armsql.ManagedDatabaseStatusOnline),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testdb2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/databases/testdb2"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Properties: &armsql.ManagedDatabaseProperties{
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-04T15:00:17.730Z"); return t}()),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				IsLedgerOn: to.Ptr(false),
		// 				Status: to.Ptr(armsql.ManagedDatabaseStatusOnline),
		// 			},
		// 	}},
		// }
	}
}
