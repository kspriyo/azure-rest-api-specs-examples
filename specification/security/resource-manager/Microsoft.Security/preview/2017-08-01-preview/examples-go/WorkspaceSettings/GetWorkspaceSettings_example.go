package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/GetWorkspaceSettings_example.json
func ExampleWorkspaceSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceSettingsClient().NewListPager(nil)
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
		// page.WorkspaceSettingList = armsecurity.WorkspaceSettingList{
		// 	Value: []*armsecurity.WorkspaceSetting{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/workspaceSettings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/workspaceSettings/default"),
		// 			Properties: &armsecurity.WorkspaceSettingProperties{
		// 				Scope: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
		// 				WorkspaceID: to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myRg"),
		// 			Type: to.Ptr("Microsoft.Security/workspaceSettings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Security/workspaceSettings/myRg"),
		// 			Properties: &armsecurity.WorkspaceSettingProperties{
		// 				Scope: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg"),
		// 				WorkspaceID: to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myOtherRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace2"),
		// 			},
		// 	}},
		// }
	}
}
