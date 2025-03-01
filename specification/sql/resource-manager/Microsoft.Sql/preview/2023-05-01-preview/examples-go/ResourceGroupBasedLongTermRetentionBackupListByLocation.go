package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupListByLocation.json
func ExampleLongTermRetentionBackupsClient_NewListByResourceGroupLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionBackupsClient().NewListByResourceGroupLocationPager("testResourceGroup", "japaneast", &armsql.LongTermRetentionBackupsClientListByResourceGroupLocationOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.LongTermRetentionBackupListResult = armsql.LongTermRetentionBackupListResult{
		// 	Value: []*armsql.LongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver1/longTermRetentionDatabases/testDatabase1/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierArchive),
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				IsBackupImmutable: to.Ptr(false),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ServerName: to.Ptr("testserver1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver1/longTermRetentionDatabases/testDatabase1/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				IsBackupImmutable: to.Ptr(false),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ServerName: to.Ptr("testserver1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12341234-1234-1234-1234-123123123123;131657960820000000;Archive"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver2/longTermRetentionDatabases/testDatabase2/longTermRetentionBackups/12341234-1234-1234-1234-123123123123;131657960820000000;Archive"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierArchive),
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase2"),
		// 				IsBackupImmutable: to.Ptr(false),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-10T08:00:00.000Z"); return t}()),
		// 				ServerName: to.Ptr("testserver2"),
		// 			},
		// 	}},
		// }
	}
}
