package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerSecurityAlertsCreateMin.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate_updateAServersThreatDetectionPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerSecurityAlertPoliciesClient().BeginCreateOrUpdate(ctx, "securityalert-4799", "securityalert-6440", armsql.SecurityAlertPolicyNameDefault, armsql.ServerSecurityAlertPolicy{
		Properties: &armsql.SecurityAlertsPolicyProperties{
			State: to.Ptr(armsql.SecurityAlertsPolicyStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerSecurityAlertPolicy = armsql.ServerSecurityAlertPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-4799/providers/Microsoft.Sql/servers/securityalert-6440/securityAlertPolicies/default"),
	// 	Properties: &armsql.SecurityAlertsPolicyProperties{
	// 		DisabledAlerts: []*string{
	// 			to.Ptr("")},
	// 			EmailAccountAdmins: to.Ptr(false),
	// 			EmailAddresses: []*string{
	// 				to.Ptr("")},
	// 				RetentionDays: to.Ptr[int32](0),
	// 				State: to.Ptr(armsql.SecurityAlertsPolicyStateEnabled),
	// 				StorageAccountAccessKey: to.Ptr(""),
	// 				StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 			},
	// 			SystemData: &armsql.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T04:41:33.937Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T04:41:33.937Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 			},
	// 		}
}
