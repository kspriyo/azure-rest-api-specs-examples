package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliances_example.json
func ExampleCompliancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCompliancesClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", nil)
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
		// page.ComplianceList = armsecurity.ComplianceList{
		// 	Value: []*armsecurity.Compliance{
		// 		{
		// 			Name: to.Ptr("2018-01-01Z"),
		// 			Type: to.Ptr("Microsoft.Security/compliances"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/compliances/2018-01-01Z"),
		// 			Properties: &armsecurity.ComplianceProperties{
		// 				AssessmentResult: []*armsecurity.ComplianceSegment{
		// 					{
		// 						Percentage: to.Ptr[float64](77.77777777777779),
		// 						SegmentType: to.Ptr("Compliant"),
		// 				}},
		// 				AssessmentTimestampUTCDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
		// 				ResourceCount: to.Ptr[int32](18),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-01-02Z"),
		// 			Type: to.Ptr("Microsoft.Security/compliances"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/compliances/2018-01-02Z"),
		// 			Properties: &armsecurity.ComplianceProperties{
		// 				AssessmentResult: []*armsecurity.ComplianceSegment{
		// 					{
		// 						Percentage: to.Ptr[float64](94.44444444444444),
		// 						SegmentType: to.Ptr("Compliant"),
		// 				}},
		// 				AssessmentTimestampUTCDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-02T00:00:00.000Z"); return t}()),
		// 				ResourceCount: to.Ptr[int32](18),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-01-03Z"),
		// 			Type: to.Ptr("Microsoft.Security/compliances"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/compliances/2018-01-03Z"),
		// 			Properties: &armsecurity.ComplianceProperties{
		// 				AssessmentResult: []*armsecurity.ComplianceSegment{
		// 					{
		// 						Percentage: to.Ptr[float64](100),
		// 						SegmentType: to.Ptr("Compliant"),
		// 				}},
		// 				AssessmentTimestampUTCDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-03T00:00:00.000Z"); return t}()),
		// 				ResourceCount: to.Ptr[int32](18),
		// 			},
		// 	}},
		// }
	}
}
