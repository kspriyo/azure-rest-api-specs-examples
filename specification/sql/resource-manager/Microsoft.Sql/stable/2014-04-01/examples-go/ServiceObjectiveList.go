package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServiceObjectiveList.json
func ExampleServiceObjectivesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceObjectivesClient().NewListByServerPager("group1", "sqlcrudtest", nil)
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
		// page.ServiceObjectiveListResult = armsql.ServiceObjectiveListResult{
		// 	Value: []*armsql.ServiceObjective{
		// 		{
		// 			Name: to.Ptr("26e021db-f1f9-4c98-84c6-92af8ef433d7"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/26e021db-f1f9-4c98-84c6-92af8ef433d7"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("29dd7459-4a7c-4e56-be22-f0adda49440d"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/29dd7459-4a7c-4e56-be22-f0adda49440d"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("c99ac918-dbea-463f-a475-16ec020fdc12"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/c99ac918-dbea-463f-a475-16ec020fdc12"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("620323bf-2879-4807-b30d-c2e6d7b3b3aa"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/620323bf-2879-4807-b30d-c2e6d7b3b3aa"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("33d0db1f-6893-4210-99f9-463fb9b496a4"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/33d0db1f-6893-4210-99f9-463fb9b496a4"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("da24338c-a6c9-46c2-a4bf-4ac95b496ae4"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/da24338c-a6c9-46c2-a4bf-4ac95b496ae4"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("53f7fa1b-b0d0-43d6-bc29-c5f059fb36e9"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/53f7fa1b-b0d0-43d6-bc29-c5f059fb36e9"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System2L"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("e79cd55c-689f-48d9-bffa-0dd12c772248"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/e79cd55c-689f-48d9-bffa-0dd12c772248"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System3L"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4b37bb6d-e004-47ac-8f7a-be56ac9fb490"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/4b37bb6d-e004-47ac-8f7a-be56ac9fb490"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(false),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(true),
		// 				ServiceObjectiveName: to.Ptr("System4L"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("6aa3bb3e-7f50-40d6-95ef-5497c30d99d8"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/6aa3bb3e-7f50-40d6-95ef-5497c30d99d8"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("Free"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dd6d99bb-f193-4ec1-86f2-43d3bccbc49c"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/dd6d99bb-f193-4ec1-86f2-43d3bccbc49c"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("Basic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("f1173c43-91bd-4aaa-973c-54e79e15235b"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/f1173c43-91bd-4aaa-973c-54e79e15235b"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("S0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1b1ebd4d-d903-4baa-97f9-4ea675f5e928"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/1b1ebd4d-d903-4baa-97f9-4ea675f5e928"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("S1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("455330e1-00cd-488b-b5fa-177c226f28b7"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/455330e1-00cd-488b-b5fa-177c226f28b7"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("S2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("789681b8-ca10-4eb0-bdf2-e0b050601b40"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/789681b8-ca10-4eb0-bdf2-e0b050601b40"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("S3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("7203483a-c4fb-4304-9e9f-17c71c904f5d"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/7203483a-c4fb-4304-9e9f-17c71c904f5d"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a7d1b92d-c987-4375-b54d-2b1d0e0f5bb0"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/a7d1b92d-c987-4375-b54d-2b1d0e0f5bb0"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("afe1eee1-1f12-4e5f-9ad6-2de9c12cb4dc"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/afe1eee1-1f12-4e5f-9ad6-2de9c12cb4dc"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("43940481-9191-475a-9dba-6b505615b9aa"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/43940481-9191-475a-9dba-6b505615b9aa"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dd00d544-bbc0-4f61-ba60-cdce0c410288"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/dd00d544-bbc0-4f61-ba60-cdce0c410288"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P11"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5bc86cca-9a96-4a94-90ef-bbdfcfbf2d71"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/5bc86cca-9a96-4a94-90ef-bbdfcfbf2d71"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("P15"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dfdc102c-ed02-4349-9756-e227f0e43bb8"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/dfdc102c-ed02-4349-9756-e227f0e43bb8"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("PRS1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a089506e-b47a-4f42-8a32-cc19af4c86fb"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/a089506e-b47a-4f42-8a32-cc19af4c86fb"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("PRS2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("39cb8faf-cba8-4b1b-b580-1e1202f2a024"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/39cb8faf-cba8-4b1b-b580-1e1202f2a024"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("PRS4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1e8da92e-efcd-4682-9140-bf6582120d1f"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/1e8da92e-efcd-4682-9140-bf6582120d1f"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("PRS6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4e63cb0e-91b9-46fd-b05c-51fdd2367618"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/4e63cb0e-91b9-46fd-b05c-51fdd2367618"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW100"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("99e78a92-d724-4e1b-857b-2be661f3d153"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/99e78a92-d724-4e1b-857b-2be661f3d153"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("284f1aff-fee7-4d3b-a211-5b8ebdd28fea"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/284f1aff-fee7-4d3b-a211-5b8ebdd28fea"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW300"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("3bdaeefe-8a9d-41d3-91c4-46ef896b19af"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/3bdaeefe-8a9d-41d3-91c4-46ef896b19af"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW400"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5f759b78-8ec0-4dfb-97cc-c1455a3b5b4d"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/5f759b78-8ec0-4dfb-97cc-c1455a3b5b4d"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW500"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("efd65c5b-af7b-4389-9109-f6a69d6a3885"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/efd65c5b-af7b-4389-9109-f6a69d6a3885"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW600"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("b89b9c6a-4ec2-4eb8-99db-6d2807e6aabb"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/b89b9c6a-4ec2-4eb8-99db-6d2807e6aabb"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW1000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("9a7a374e-b95c-4fd5-a68e-131d60796c47"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/9a7a374e-b95c-4fd5-a68e-131d60796c47"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW1200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("b930f58e-86b5-43e0-a2da-d8bf8769c557"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/b930f58e-86b5-43e0-a2da-d8bf8769c557"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW1500"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("99165ede-a5ab-4b52-b317-e391d92ec370"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/99165ede-a5ab-4b52-b317-e391d92ec370"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW2000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("8e28c923-5cf2-43cb-bd25-28c8c69b30ff"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/8e28c923-5cf2-43cb-bd25-28c8c69b30ff"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW3000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ee1df062-4f3c-42ad-91bf-58b2a7c351e4"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/ee1df062-4f3c-42ad-91bf-58b2a7c351e4"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DW6000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("9cfc850f-d57f-4760-b5a6-bb640d268bf0"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/9cfc850f-d57f-4760-b5a6-bb640d268bf0"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(true),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS100"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("053407ef-f01c-46f4-b829-96e01a14f449"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/053407ef-f01c-46f4-b829-96e01a14f449"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("013a9e10-cafc-45a8-8fcf-93095655d2ce"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/013a9e10-cafc-45a8-8fcf-93095655d2ce"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS300"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("79f61db4-8c10-46ba-a93a-d7d02dddd61c"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/79f61db4-8c10-46ba-a93a-d7d02dddd61c"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS400"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("44eaac33-df00-4ef4-a2bb-f7ff87899eea"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/44eaac33-df00-4ef4-a2bb-f7ff87899eea"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS500"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("f8e0f3a6-888b-459c-a9dd-d74d8b2b0e72"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/f8e0f3a6-888b-459c-a9dd-d74d8b2b0e72"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS600"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("b9ed8f51-a414-42dc-8348-e4a1de25e12b"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/b9ed8f51-a414-42dc-8348-e4a1de25e12b"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS1000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("07479569-6d70-47a5-8db6-0af55d34f2c1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/07479569-6d70-47a5-8db6-0af55d34f2c1"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS1200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2d79baec-2879-46d5-9f5d-fb70eb004c4e"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/2d79baec-2879-46d5-9f5d-fb70eb004c4e"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS1500"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("7fb5389f-6d15-4e0b-9540-fe5ecdfdbeee"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/7fb5389f-6d15-4e0b-9540-fe5ecdfdbeee"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("DS2000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d1737d22-a8ea-4de7-9bd0-33395d2a7419"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlcrudtest/serviceObjectives/d1737d22-a8ea-4de7-9bd0-33395d2a7419"),
		// 			Properties: &armsql.ServiceObjectiveProperties{
		// 				Enabled: to.Ptr(true),
		// 				IsDefault: to.Ptr(false),
		// 				IsSystem: to.Ptr(false),
		// 				ServiceObjectiveName: to.Ptr("ElasticPool"),
		// 			},
		// 	}},
		// }
	}
}
