package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ListRestorableDroppedDatabasesByServer.json
func ExampleRestorableDroppedDatabasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableDroppedDatabasesClient().NewListByServerPager("Default-SQL-SouthEastAsia", "testsvr", nil)
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
		// page.RestorableDroppedDatabaseListResult = armsql.RestorableDroppedDatabaseListResult{
		// 	Value: []*armsql.RestorableDroppedDatabase{
		// 		{
		// 			Name: to.Ptr("testdb,131403269876900000"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/restorableDroppedDatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases/testdb"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Properties: &armsql.RestorableDroppedDatabaseProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 				DatabaseName: to.Ptr("testdb"),
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T02:49:47.690Z"); return t}()),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("BC_Gen4_2"),
		// 				Tier: to.Ptr("BusinessCritical"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testdb2,131403269876900000"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/restorableDroppedDatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases/testdb2"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Properties: &armsql.RestorableDroppedDatabaseProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 				DatabaseName: to.Ptr("testdb2"),
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T02:49:47.690Z"); return t}()),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 	}},
		// }
	}
}
