package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceStandardList_example.json
func ExampleRegulatoryComplianceStandardsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegulatoryComplianceStandardsClient().NewListPager(&armsecurity.RegulatoryComplianceStandardsClientListOptions{Filter: nil})
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
		// page.RegulatoryComplianceStandardList = armsecurity.RegulatoryComplianceStandardList{
		// 	Value: []*armsecurity.RegulatoryComplianceStandard{
		// 		{
		// 			Name: to.Ptr("PCI-DSS-3.2"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceStandard"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2"),
		// 			Properties: &armsecurity.RegulatoryComplianceStandardProperties{
		// 				FailedControls: to.Ptr[int32](4),
		// 				PassedControls: to.Ptr[int32](7),
		// 				SkippedControls: to.Ptr[int32](0),
		// 				State: to.Ptr(armsecurity.StateFailed),
		// 				UnsupportedControls: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ISO-27001"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceStandard"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/ISO-27001"),
		// 			Properties: &armsecurity.RegulatoryComplianceStandardProperties{
		// 				FailedControls: to.Ptr[int32](0),
		// 				PassedControls: to.Ptr[int32](0),
		// 				SkippedControls: to.Ptr[int32](10),
		// 				State: to.Ptr(armsecurity.StateSkipped),
		// 				UnsupportedControls: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AZURE-CIS"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceStandard"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/AZURE-CIS"),
		// 			Properties: &armsecurity.RegulatoryComplianceStandardProperties{
		// 				FailedControls: to.Ptr[int32](0),
		// 				PassedControls: to.Ptr[int32](0),
		// 				SkippedControls: to.Ptr[int32](0),
		// 				State: to.Ptr(armsecurity.StateUnsupported),
		// 				UnsupportedControls: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SOC-TSP"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceStandard"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/SOC-TSP"),
		// 			Properties: &armsecurity.RegulatoryComplianceStandardProperties{
		// 				FailedControls: to.Ptr[int32](0),
		// 				PassedControls: to.Ptr[int32](15),
		// 				SkippedControls: to.Ptr[int32](0),
		// 				State: to.Ptr(armsecurity.StatePassed),
		// 				UnsupportedControls: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
