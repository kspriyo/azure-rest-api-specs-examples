package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliance_example.json
func ExampleCompliancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCompliancesClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "2018-01-01Z", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Compliance = armsecurity.Compliance{
	// 	Name: to.Ptr("2018-01-01Z"),
	// 	Type: to.Ptr("Microsoft.Security/compliances"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/compliances/2018-01-01Z"),
	// 	Properties: &armsecurity.ComplianceProperties{
	// 		AssessmentResult: []*armsecurity.ComplianceSegment{
	// 			{
	// 				Percentage: to.Ptr[float64](77.77777777777779),
	// 				SegmentType: to.Ptr("Compliant"),
	// 		}},
	// 		AssessmentTimestampUTCDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 		ResourceCount: to.Ptr[int32](18),
	// 	},
	// }
}
