package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTasksSubscription_example.json
func ExampleTasksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTasksClient().NewListPager(&armsecurity.TasksClientListOptions{Filter: nil})
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
		// page.TaskList = armsecurity.TaskList{
		// 	Value: []*armsecurity.Task{
		// 		{
		// 			Name: to.Ptr("62609ee7-d0a5-8616-9fe4-1df5cca7758d"),
		// 			Type: to.Ptr("Microsoft.Security/locations/tasks"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/locations/westeurope/tasks/62609ee7-d0a5-8616-9fe4-1df5cca7758d"),
		// 			Properties: &armsecurity.TaskProperties{
		// 				CreationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-05T10:42:03.993Z"); return t}()),
		// 				LastStateChangeTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-05T10:42:03.993Z"); return t}()),
		// 				SecurityTaskParameters: &armsecurity.TaskParameters{
		// 					AdditionalProperties: map[string]any{
		// 						"location": "uksouth",
		// 						"resourceGroup": "myRg",
		// 						"resourceId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/default",
		// 						"resourceName": "default",
		// 						"resourceParent": "vnet1",
		// 						"resourceType": "Subnet",
		// 						"uniqueKey": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/default",
		// 					},
		// 					Name: to.Ptr("NetworkSecurityGroupMissingOnSubnet"),
		// 				},
		// 				State: to.Ptr("Active"),
		// 				SubState: to.Ptr("NA"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d55b4dc0-779c-c66c-33e5-d7bce24c4222"),
		// 			Type: to.Ptr("Microsoft.Security/locations/tasks"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Security/locations/westeurope/tasks/d55b4dc0-779c-c66c-33e5-d7bce24c4222"),
		// 			Properties: &armsecurity.TaskProperties{
		// 				CreationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-02T11:41:27.054Z"); return t}()),
		// 				LastStateChangeTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-02T11:41:27.054Z"); return t}()),
		// 				SecurityTaskParameters: &armsecurity.TaskParameters{
		// 					AdditionalProperties: map[string]any{
		// 						"isDataDiskEncrypted": false,
		// 						"isOsDiskEncrypted": false,
		// 						"resourceId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
		// 						"severity": "High",
		// 						"uniqueKey": "EncryptionOnVmTaskParameters_/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
		// 						"vmId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
		// 						"vmName": "vm1",
		// 					},
		// 					Name: to.Ptr("EncryptionOnVm"),
		// 				},
		// 				State: to.Ptr("Active"),
		// 				SubState: to.Ptr("NA"),
		// 			},
		// 	}},
		// }
	}
}
