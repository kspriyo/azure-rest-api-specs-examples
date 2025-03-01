package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/ListComplianceResults_example.json
func ExampleComplianceResultsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComplianceResultsClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", nil)
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
		// page.ComplianceResultList = armsecurity.ComplianceResultList{
		// 	Value: []*armsecurity.ComplianceResult{
		// 		{
		// 			Name: to.Ptr("DesignateMoreThanOneOwner"),
		// 			Type: to.Ptr("Microsoft.Security/complianceResults"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/complianceResults/DesignateMoreThanOneOwner"),
		// 			Properties: &armsecurity.ComplianceResultProperties{
		// 				ResourceStatus: to.Ptr(armsecurity.ResourceStatusHealthy),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RemoveExternalAccountsWithReadPermissions"),
		// 			Type: to.Ptr("Microsoft.Security/complianceResults"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/complianceResults/RemoveExternalAccountsWithReadPermissions"),
		// 			Properties: &armsecurity.ComplianceResultProperties{
		// 				ResourceStatus: to.Ptr(armsecurity.ResourceStatusNotHealthy),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RemoveDeprecatedAccounts"),
		// 			Type: to.Ptr("Microsoft.Security/complianceResults"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/complianceResults/RemoveDeprecatedAccounts"),
		// 			Properties: &armsecurity.ComplianceResultProperties{
		// 				ResourceStatus: to.Ptr(armsecurity.ResourceStatusHealthy),
		// 			},
		// 	}},
		// }
	}
}
