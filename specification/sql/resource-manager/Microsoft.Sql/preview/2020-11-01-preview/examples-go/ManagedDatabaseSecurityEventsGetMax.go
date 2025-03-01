package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityEventsGetMax.json
func ExampleManagedDatabaseSecurityEventsClient_NewListByDatabasePager_getTheManagedDatabasesSecurityEventsWithMaximalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseSecurityEventsClient().NewListByDatabasePager("testrg", "testcl", "database1", &armsql.ManagedDatabaseSecurityEventsClientListByDatabaseOptions{Filter: to.Ptr("ShowServerRecords eq true"),
		Skip:      to.Ptr[int32](0),
		Top:       to.Ptr[int32](1),
		Skiptoken: to.Ptr("eyJCbG9iTmFtZURhdGVUaW1lIjoiXC9EYXRlKDE1MTIyODg4MTIwMTArMDIwMClcLyIsIkJsb2JOYW1lUm9sbG92ZXJJbmRleCI6IjAiLCJFbmREYXRlIjoiXC9EYXRlKDE1MTI0NjYyMDA1MjkpXC8iLCJJc1NraXBUb2tlblNldCI6ZmFsc2UsIklzVjJCbG9iVGltZUZvcm1hdCI6dHJ1ZSwiU2hvd1NlcnZlclJlY29yZHMiOmZhbHNlLCJTa2lwVmFsdWUiOjAsIlRha2VWYWx1ZSI6MTB9"),
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
		// page.SecurityEventCollection = armsql.SecurityEventCollection{
		// 	Value: []*armsql.SecurityEvent{
		// 		{
		// 			Name: to.Ptr("06364798761800000000000000001"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/securityEvents"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/databases/database1/securityEvents/06364798761800000000000000001"),
		// 			Properties: &armsql.SecurityEventProperties{
		// 				ApplicationName: to.Ptr("myApp"),
		// 				ClientIP: to.Ptr("10.166.113.220"),
		// 				Database: to.Ptr("database1"),
		// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-24T10:13:24.729Z"); return t}()),
		// 				PrincipalName: to.Ptr("maliciousUser"),
		// 				SecurityEventSQLInjectionAdditionalProperties: &armsql.SecurityEventSQLInjectionAdditionalProperties{
		// 					ErrorCode: to.Ptr[int32](0),
		// 					ErrorMessage: to.Ptr(""),
		// 					ErrorSeverity: to.Ptr[int32](0),
		// 					Statement: to.Ptr("select * from sys.databases where database_id like '' or 1 = 1 --' and family = 'test11'"),
		// 					StatementHighlightLength: to.Ptr[int32](13),
		// 					StatementHighlightOffset: to.Ptr[int32](52),
		// 					ThreatID: to.Ptr("1"),
		// 				},
		// 				SecurityEventType: to.Ptr(armsql.SecurityEventTypeSQLInjectionExploit),
		// 				Server: to.Ptr("testcl"),
		// 				Subscription: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 			},
		// 	}},
		// }
	}
}
