package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListRecoverableManagedDatabasesByServer.json
func ExampleRecoverableManagedDatabasesClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecoverableManagedDatabasesClient().NewListByInstancePager("Test1", "managedInstance", nil)
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
		// page.RecoverableManagedDatabaseListResult = armsql.RecoverableManagedDatabaseListResult{
		// 	Value: []*armsql.RecoverableManagedDatabase{
		// 		{
		// 			Name: to.Ptr("testdb"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/recoverableDaatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/recoverableDaatabases/testdb"),
		// 			Properties: &armsql.RecoverableManagedDatabaseProperties{
		// 				LastAvailableBackupDate: to.Ptr("2017-05-26T01:06:29.78Z"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testdb2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/recoverableDaatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/recoverableDaatabases/testdb2"),
		// 			Properties: &armsql.RecoverableManagedDatabaseProperties{
		// 				LastAvailableBackupDate: to.Ptr("2018-04-26T02:17:23.44Z"),
		// 			},
		// 	}},
		// }
	}
}
