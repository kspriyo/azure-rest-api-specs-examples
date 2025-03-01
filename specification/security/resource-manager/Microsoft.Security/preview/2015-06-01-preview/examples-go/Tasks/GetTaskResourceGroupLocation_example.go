package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTaskResourceGroupLocation_example.json
func ExampleTasksClient_GetResourceGroupLevelTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().GetResourceGroupLevelTask(ctx, "myRg", "westeurope", "d55b4dc0-779c-c66c-33e5-d7bce24c4222", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Task = armsecurity.Task{
	// 	Name: to.Ptr("d55b4dc0-779c-c66c-33e5-d7bce24c4222"),
	// 	Type: to.Ptr("Microsoft.Security/locations/tasks"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Security/locations/westeurope/tasks/d55b4dc0-779c-c66c-33e5-d7bce24c4222"),
	// 	Properties: &armsecurity.TaskProperties{
	// 		CreationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-02T11:41:27.054Z"); return t}()),
	// 		LastStateChangeTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-02T11:41:27.054Z"); return t}()),
	// 		SecurityTaskParameters: &armsecurity.TaskParameters{
	// 			AdditionalProperties: map[string]any{
	// 				"isDataDiskEncrypted": false,
	// 				"isOsDiskEncrypted": false,
	// 				"resourceId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
	// 				"severity": "High",
	// 				"uniqueKey": "EncryptionOnVmTaskParameters_/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
	// 				"vmId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
	// 				"vmName": "vm1",
	// 			},
	// 			Name: to.Ptr("EncryptionOnVm"),
	// 		},
	// 		State: to.Ptr("Active"),
	// 		SubState: to.Ptr("NA"),
	// 	},
	// }
}
