package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutions/GetSecuritySolutionsSubscription_example.json
func ExampleSolutionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSolutionsClient().NewListPager(nil)
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
		// page.SolutionList = armsecurity.SolutionList{
		// 	Value: []*armsecurity.Solution{
		// 		{
		// 			Location: to.Ptr("eastus"),
		// 			Name: to.Ptr("MySaasWaf"),
		// 			Type: to.Ptr("Microsoft.Security/locations/securitySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/centralus/securitySolutions/SaasWaf"),
		// 			Properties: &armsecurity.SolutionProperties{
		// 				ProtectionStatus: to.Ptr("Good"),
		// 				ProvisioningState: to.Ptr(armsecurity.ProvisioningStateSucceeded),
		// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilySaasWaf),
		// 				Template: to.Ptr("microsoft/ApplicationGateway-ARM"),
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("eastus2"),
		// 			Name: to.Ptr("MyVA"),
		// 			Type: to.Ptr("Microsoft.Security/locations/securitySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg2/providers/Microsoft.Security/locations/centralus/securitySolutions/MyVA"),
		// 			Properties: &armsecurity.SolutionProperties{
		// 				ProtectionStatus: to.Ptr("Good"),
		// 				ProvisioningState: to.Ptr(armsecurity.ProvisioningStateSucceeded),
		// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilyVa),
		// 				Template: to.Ptr("qualys.qualysAgent"),
		// 			},
		// 	}},
		// }
	}
}
