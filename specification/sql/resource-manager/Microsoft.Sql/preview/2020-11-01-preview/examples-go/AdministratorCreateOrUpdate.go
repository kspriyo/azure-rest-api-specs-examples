package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/AdministratorCreateOrUpdate.json
func ExampleServerAzureADAdministratorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerAzureADAdministratorsClient().BeginCreateOrUpdate(ctx, "sqlcrudtest-4799", "sqlcrudtest-6440", armsql.AdministratorNameActiveDirectory, armsql.ServerAzureADAdministrator{
		Properties: &armsql.AdministratorProperties{
			AdministratorType: to.Ptr(armsql.AdministratorTypeActiveDirectory),
			Login:             to.Ptr("bob@contoso.com"),
			Sid:               to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			TenantID:          to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
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
	// res.ServerAzureADAdministrator = armsql.ServerAzureADAdministrator{
	// 	Name: to.Ptr("ActiveDirectory"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/administrators"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-6440/administrators/ActiveDirectory"),
	// 	Properties: &armsql.AdministratorProperties{
	// 		AdministratorType: to.Ptr(armsql.AdministratorTypeActiveDirectory),
	// 		AzureADOnlyAuthentication: to.Ptr(true),
	// 		Login: to.Ptr("bob@contoso.com"),
	// 		Sid: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 		TenantID: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 	},
	// }
}
