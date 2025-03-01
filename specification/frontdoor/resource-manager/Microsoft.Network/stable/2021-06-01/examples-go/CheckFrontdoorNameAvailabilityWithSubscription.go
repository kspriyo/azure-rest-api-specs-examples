package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/CheckFrontdoorNameAvailabilityWithSubscription.json
func ExampleNameAvailabilityWithSubscriptionClient_Check() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNameAvailabilityWithSubscriptionClient().Check(ctx, armfrontdoor.CheckNameAvailabilityInput{
		Name: to.Ptr("sampleName"),
		Type: to.Ptr(armfrontdoor.ResourceTypeMicrosoftNetworkFrontDoorsFrontendEndpoints),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityOutput = armfrontdoor.CheckNameAvailabilityOutput{
	// 	Message: to.Ptr("Name not available"),
	// 	NameAvailability: to.Ptr(armfrontdoor.AvailabilityUnavailable),
	// 	Reason: to.Ptr("Name is already in use"),
	// }
}
