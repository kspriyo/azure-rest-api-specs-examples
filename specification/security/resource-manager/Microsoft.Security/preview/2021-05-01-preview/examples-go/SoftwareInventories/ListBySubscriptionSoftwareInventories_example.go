package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/ListBySubscriptionSoftwareInventories_example.json
func ExampleSoftwareInventoriesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSoftwareInventoriesClient().NewListBySubscriptionPager(nil)
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
		// page.SoftwaresList = armsecurity.SoftwaresList{
		// 	Value: []*armsecurity.Software{
		// 		{
		// 			Name: to.Ptr("outlook_16.0.10371.20060"),
		// 			Type: to.Ptr("Microsoft.Security/softwareInventories"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/EITAN-TESTS/providers/Microsoft.Compute/virtualMachines/Eitan-Test1/providers/Microsoft.Security/softwareInventories/outlook_16.0.10371.20060"),
		// 			Properties: &armsecurity.SoftwareProperties{
		// 				DeviceID: to.Ptr("7bd19ed6b07553e52a2844451bcec68d25963a53"),
		// 				EndOfSupportStatus: to.Ptr(armsecurity.EndOfSupportStatusNone),
		// 				FirstSeenAt: to.Ptr("2021-01-26 15:48:56"),
		// 				NumberOfKnownVulnerabilities: to.Ptr[int32](0),
		// 				OSPlatform: to.Ptr("Windows10"),
		// 				SoftwareName: to.Ptr("outlook"),
		// 				Vendor: to.Ptr("microsoft"),
		// 				Version: to.Ptr("16.0.10371.20060"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("windows_10_10.0.19042.746"),
		// 			Type: to.Ptr("Microsoft.Security/softwareInventories"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/EITAN-TESTS/providers/Microsoft.Compute/virtualMachines/Eitan-Test2/providers/Microsoft.Security/softwareInventories/windows_10_10.0.19042.746"),
		// 			Properties: &armsecurity.SoftwareProperties{
		// 				DeviceID: to.Ptr("7bd19ed6b07553e52a2844451bcec68d25963a53"),
		// 				EndOfSupportStatus: to.Ptr(armsecurity.EndOfSupportStatusNone),
		// 				FirstSeenAt: to.Ptr("2021-01-26 15:51:19"),
		// 				NumberOfKnownVulnerabilities: to.Ptr[int32](26),
		// 				OSPlatform: to.Ptr("Windows10"),
		// 				SoftwareName: to.Ptr("windows_10"),
		// 				Vendor: to.Ptr("microsoft"),
		// 				Version: to.Ptr("10.0.19042.746"),
		// 			},
		// 	}},
		// }
	}
}
