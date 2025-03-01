package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/AdministratorList.json
func ExampleServerAzureADAdministratorsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerAzureADAdministratorsClient().NewListByServerPager("sqlcrudtest-4799", "sqlcrudtest-6440", nil)
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
		// page.AdministratorListResult = armsql.AdministratorListResult{
		// 	Value: []*armsql.ServerAzureADAdministrator{
		// 		{
		// 			Name: to.Ptr("ActiveDirectory"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/administrators"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-6440/administrators/ActiveDirectory"),
		// 			Properties: &armsql.AdministratorProperties{
		// 				AdministratorType: to.Ptr(armsql.AdministratorTypeActiveDirectory),
		// 				AzureADOnlyAuthentication: to.Ptr(true),
		// 				Login: to.Ptr("bob@contoso.com"),
		// 				Sid: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
		// 				TenantID: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
		// 			},
		// 	}},
		// }
	}
}
