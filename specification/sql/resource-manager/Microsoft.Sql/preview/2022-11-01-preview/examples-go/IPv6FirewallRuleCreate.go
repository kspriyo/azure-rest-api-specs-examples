package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleCreate.json
func ExampleIPv6FirewallRulesClient_CreateOrUpdate_createAnIPv6FirewallRuleMaxMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPv6FirewallRulesClient().CreateOrUpdate(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-5370", armsql.IPv6FirewallRule{
		Properties: &armsql.IPv6ServerFirewallRuleProperties{
			EndIPv6Address:   to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
			StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPv6FirewallRule = armsql.IPv6FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-5370"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5370"),
	// 	Properties: &armsql.IPv6ServerFirewallRuleProperties{
	// 		EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
	// 		StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
	// 	},
	// }
}
