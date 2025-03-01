package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/officeConsents/GetOfficeConsentsById.json
func ExampleOfficeConsentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOfficeConsentsClient().Get(ctx, "myRg", "myWorkspace", "04e5fd05-ff86-4b97-b8d2-1c20933cb46c", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OfficeConsent = armsecurityinsights.OfficeConsent{
	// 	Name: to.Ptr("04e5fd05-ff86-4b97-b8d2-1c20933cb46c"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/officeConsents"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/officeConsents/04e5fd05-ff86-4b97-b8d2-1c20933cb46c"),
	// 	Properties: &armsecurityinsights.OfficeConsentProperties{
	// 		ConsentID: to.Ptr("04e5fd05-ff86-4b97-b8d2-1c20933cb46c"),
	// 		TenantID: to.Ptr("5460b3d2-1e7b-4757-ad54-c858c7e3f252"),
	// 	},
	// }
}
