package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutionsReferenceData/GetSecuritySolutionsReferenceDataSubscriptionLocation_example.json
func ExampleSolutionsReferenceDataClient_ListByHomeRegion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionsReferenceDataClient().ListByHomeRegion(ctx, "westcentralus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SolutionsReferenceDataList = armsecurity.SolutionsReferenceDataList{
	// 	Value: []*armsecurity.SolutionsReferenceData{
	// 		{
	// 			Name: to.Ptr("microsoft.ApplicationGateway-ARM"),
	// 			Type: to.Ptr("Microsoft.Security/locations/securitySolutionsReferenceData"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/locations/westcentralus/securitySolutionsReferenceData/microsoft.ApplicationGateway-ARM"),
	// 			Properties: &armsecurity.SolutionsReferenceDataProperties{
	// 				AlertVendorName: to.Ptr("Microsoft"),
	// 				PackageInfoURL: to.Ptr("www.azure.com"),
	// 				ProductName: to.Ptr("Web Application Firewall"),
	// 				Publisher: to.Ptr("microsoft"),
	// 				PublisherDisplayName: to.Ptr("Microsoft Inc."),
	// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilySaasWaf),
	// 				Template: to.Ptr("microsoft/ApplicationGateway-ARM"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("qualys.qualysAgent"),
	// 			Type: to.Ptr("Microsoft.Security/locations/SecuritySolutions"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/locations/westcentralus/securitySolutionsReferenceData/qualys.qualysAgent"),
	// 			Properties: &armsecurity.SolutionsReferenceDataProperties{
	// 				AlertVendorName: to.Ptr("Qualys VA"),
	// 				PackageInfoURL: to.Ptr("http://www.qualys.com/"),
	// 				ProductName: to.Ptr("Vulnerability Assessment"),
	// 				Publisher: to.Ptr("qualys"),
	// 				PublisherDisplayName: to.Ptr("Qualys, Inc."),
	// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilyVa),
	// 				Template: to.Ptr("qualys/qualysAgent"),
	// 			},
	// 	}},
	// }
}
