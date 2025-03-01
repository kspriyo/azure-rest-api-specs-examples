package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetManagedShortTermRetentionPolicyRestorableDropped.json
func ExampleManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", armsql.ManagedShortTermRetentionPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedBackupShortTermRetentionPolicy = armsql.ManagedBackupShortTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/restorableDroppedDatabases/backupShortTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstance/testsvr/restorableDroppedDatabases/testdb,131403269876900000/backupShortTermRetentionPolicies/default"),
	// 	Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
	// 		RetentionDays: to.Ptr[int32](14),
	// 	},
	// }
}
