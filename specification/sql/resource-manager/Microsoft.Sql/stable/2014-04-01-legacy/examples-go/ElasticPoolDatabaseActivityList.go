package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ElasticPoolDatabaseActivityList.json
func ExampleElasticPoolDatabaseActivitiesClient_NewListByElasticPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticPoolDatabaseActivitiesClient().NewListByElasticPoolPager("sqlcrudtest-4673", "sqlcrudtest-603", "7537", nil)
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
		// page.ElasticPoolDatabaseActivityListResult = armsql.ElasticPoolDatabaseActivityListResult{
		// 	Value: []*armsql.ElasticPoolDatabaseActivity{
		// 		{
		// 			Name: to.Ptr("3a3272b3-f1fe-423c-9feb-7b843157eda5"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/elasticPools/elasticPoolDatabaseActivity"),
		// 			ID: to.Ptr("/subscriptions/9d4e2ad0-e20b-4464-9219-353bded52513/resourceGroups/sqlcrudtest-4673/providers/Microsoft.Sql/servers/sqlcrudtest-603/elasticPools/7537/elasticPoolDatabaseActivity/3a3272b3-f1fe-423c-9feb-7b843157eda5"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.ElasticPoolDatabaseActivityProperties{
		// 				OperationID: to.Ptr("3a3272b3-f1fe-423c-9feb-7b843157eda5"),
		// 				CurrentElasticPoolName: to.Ptr("7537"),
		// 				CurrentServiceObjective: to.Ptr("ElasticPool"),
		// 				DatabaseName: to.Ptr("2396"),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-23T03:08:02.950Z"); return t}()),
		// 				Operation: to.Ptr("UPDATE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("sqlcrudtest-603"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-23T03:07:12.570Z"); return t}()),
		// 				State: to.Ptr("COMPLETED"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("95108a78-384e-48d3-b4de-7bf23b93a26d"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/elasticPools/elasticPoolDatabaseActivity"),
		// 			ID: to.Ptr("/subscriptions/9d4e2ad0-e20b-4464-9219-353bded52513/resourceGroups/sqlcrudtest-4673/providers/Microsoft.Sql/servers/sqlcrudtest-603/elasticPools/7537/elasticPoolDatabaseActivity/95108a78-384e-48d3-b4de-7bf23b93a26d"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.ElasticPoolDatabaseActivityProperties{
		// 				OperationID: to.Ptr("95108a78-384e-48d3-b4de-7bf23b93a26d"),
		// 				CurrentElasticPoolName: to.Ptr("7537"),
		// 				CurrentServiceObjective: to.Ptr("ElasticPool"),
		// 				DatabaseName: to.Ptr("2396"),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-23T03:06:49.190Z"); return t}()),
		// 				Operation: to.Ptr("CREATE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("sqlcrudtest-603"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-23T03:06:11.190Z"); return t}()),
		// 				State: to.Ptr("COMPLETED"),
		// 			},
		// 	}},
		// }
	}
}
