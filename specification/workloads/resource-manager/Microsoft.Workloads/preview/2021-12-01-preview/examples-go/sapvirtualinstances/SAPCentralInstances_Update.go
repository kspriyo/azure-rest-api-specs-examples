package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPCentralInstances_Update.json
func ExampleSAPCentralInstancesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewSAPCentralInstancesClient("6d875e77-e412-4d7d-9af4-8895278b4443", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"test-rg",
		"X00",
		"centralServer",
		armworkloads.UpdateSAPCentralInstanceRequest{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
