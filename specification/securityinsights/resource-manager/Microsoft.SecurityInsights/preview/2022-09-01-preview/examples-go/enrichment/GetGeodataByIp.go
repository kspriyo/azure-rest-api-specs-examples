package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/enrichment/GetGeodataByIp.json
func ExampleIPGeodataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPGeodataClient().Get(ctx, "myRg", "1.2.3.4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnrichmentIPGeodata = armsecurityinsights.EnrichmentIPGeodata{
	// 	Asn: to.Ptr("12345"),
	// 	Carrier: to.Ptr("Microsoft"),
	// 	City: to.Ptr("Redmond"),
	// 	CityCf: to.Ptr[int32](90),
	// 	Continent: to.Ptr("north america"),
	// 	Country: to.Ptr("united states"),
	// 	CountryCf: to.Ptr[int32](99),
	// 	IPAddr: to.Ptr("1.2.3.4"),
	// 	IPRoutingType: to.Ptr("fixed"),
	// 	Latitude: to.Ptr("40.2436"),
	// 	Longitude: to.Ptr("-100.8891"),
	// 	Organization: to.Ptr("Microsoft"),
	// 	OrganizationType: to.Ptr("tech"),
	// 	Region: to.Ptr("western usa"),
	// 	State: to.Ptr("washington"),
	// 	StateCode: to.Ptr("wa"),
	// }
}
