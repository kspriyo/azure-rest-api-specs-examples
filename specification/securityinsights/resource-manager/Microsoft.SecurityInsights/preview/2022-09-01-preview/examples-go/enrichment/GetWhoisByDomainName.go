package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/enrichment/GetWhoisByDomainName.json
func ExampleDomainWhoisClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDomainWhoisClient().Get(ctx, "myRg", "microsoft.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnrichmentDomainWhois = armsecurityinsights.EnrichmentDomainWhois{
	// 	Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T16:15:01.187045Z"); return t}()),
	// 	Domain: to.Ptr("microsoft.com"),
	// 	ParsedWhois: &armsecurityinsights.EnrichmentDomainWhoisDetails{
	// 		Contacts: &armsecurityinsights.EnrichmentDomainWhoisContacts{
	// 			Admin: &armsecurityinsights.EnrichmentDomainWhoisContact{
	// 				Name: to.Ptr("Administrator"),
	// 				Country: to.Ptr("United States"),
	// 				Email: to.Ptr("mail@microsoft.com"),
	// 				Org: to.Ptr("Microsoft"),
	// 				Phone: to.Ptr("1-800-555-1234"),
	// 				Postal: to.Ptr("98052"),
	// 				State: to.Ptr("WA"),
	// 				Street: []*string{
	// 					to.Ptr("One Microsoft Way")},
	// 				},
	// 				Billing: &armsecurityinsights.EnrichmentDomainWhoisContact{
	// 					Name: to.Ptr("Administrator"),
	// 					Country: to.Ptr("United States"),
	// 					Email: to.Ptr("mail@microsoft.com"),
	// 					Org: to.Ptr("Microsoft"),
	// 					Phone: to.Ptr("1-800-555-1234"),
	// 					Postal: to.Ptr("98052"),
	// 					State: to.Ptr("WA"),
	// 					Street: []*string{
	// 						to.Ptr("One Microsoft Way")},
	// 					},
	// 					Tech: &armsecurityinsights.EnrichmentDomainWhoisContact{
	// 						Name: to.Ptr("Administrator"),
	// 						Country: to.Ptr("United States"),
	// 						Email: to.Ptr("mail@microsoft.com"),
	// 						Org: to.Ptr("Microsoft"),
	// 						Phone: to.Ptr("1-800-555-1234"),
	// 						Postal: to.Ptr("98052"),
	// 						State: to.Ptr("WA"),
	// 						Street: []*string{
	// 							to.Ptr("One Microsoft Way")},
	// 						},
	// 					},
	// 					NameServers: []*string{
	// 						to.Ptr("ns1-205.azure-dns.com"),
	// 						to.Ptr("ns2-205.azure-dns.net"),
	// 						to.Ptr("ns3-205.azure-dns.org"),
	// 						to.Ptr("ns4-205.azure-dns.info")},
	// 						Registrar: &armsecurityinsights.EnrichmentDomainWhoisRegistrarDetails{
	// 							Name: to.Ptr("MarkMonitor, Inc"),
	// 							AbuseContactEmail: to.Ptr("abuse@microsoft.com"),
	// 							AbuseContactPhone: to.Ptr("12083895770"),
	// 							URL: to.Ptr("http://www.markmonitor.com"),
	// 							WhoisServer: to.Ptr("whois.markmonitor.com"),
	// 						},
	// 						Statuses: []*string{
	// 							to.Ptr("clientUpdateProhibited"),
	// 							to.Ptr("clientTransferProhibited"),
	// 							to.Ptr("clientDeleteProhibited"),
	// 							to.Ptr("serverUpdateProhibited"),
	// 							to.Ptr("serverTransferProhibited"),
	// 							to.Ptr("serverDeleteProhibited")},
	// 						},
	// 						Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T16:15:01.187045Z"); return t}()),
	// 					}
}
