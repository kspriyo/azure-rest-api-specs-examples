package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ElasticPoolActivityList.json
func ExampleElasticPoolActivitiesClient_NewListByElasticPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticPoolActivitiesClient().NewListByElasticPoolPager("sqlcrudtest-4291", "sqlcrudtest-6574", "8749", nil)
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
		// page.ElasticPoolActivityListResult = armsql.ElasticPoolActivityListResult{
		// 	Value: []*armsql.ElasticPoolActivity{
		// 		{
		// 			Name: to.Ptr("851f1672-f7f0-46f6-a262-ee9b51e18e97"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/elasticPools/elasticPoolActivity"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4291/providers/Microsoft.Sql/servers/sqlcrudtest-6574/elasticPools/8749/elasticPoolActivity/851f1672-f7f0-46f6-a262-ee9b51e18e97"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.ElasticPoolActivityProperties{
		// 				OperationID: to.Ptr("851f1672-f7f0-46f6-a262-ee9b51e18e97"),
		// 				ElasticPoolName: to.Ptr("8749"),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T02:03:23.263Z"); return t}()),
		// 				Operation: to.Ptr("CREATE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				RequestedDatabaseDtuCap: to.Ptr[int32](5),
		// 				RequestedDatabaseDtuGuarantee: to.Ptr[int32](0),
		// 				RequestedDtuGuarantee: to.Ptr[int32](100),
		// 				RequestedStorageLimitInGB: to.Ptr[int64](9),
		// 				RequestedStorageLimitInMB: to.Ptr[int32](10000),
		// 				ServerName: to.Ptr("sqlcrudtest-6574"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T02:03:06.160Z"); return t}()),
		// 				State: to.Ptr("COMPLETED"),
		// 			},
		// 	}},
		// }
	}
}
