package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2022-12-01-preview/examples/DefenderForStorage/GetDefenderForStorageSettings_example.json
func ExampleDefenderForStorageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefenderForStorageClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount", armsecurity.SettingNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefenderForStorageSetting = armsecurity.DefenderForStorageSetting{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Security/defenderForStorageSettings"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount/providers/Microsoft.Security/defenderForStorageSettings/current"),
	// 	Properties: &armsecurity.DefenderForStorageSettingProperties{
	// 		IsEnabled: to.Ptr(true),
	// 		MalwareScanning: &armsecurity.MalwareScanningProperties{
	// 			OnUpload: &armsecurity.OnUploadProperties{
	// 				CapGBPerMonth: to.Ptr[int32](-1),
	// 				IsEnabled: to.Ptr(true),
	// 			},
	// 			ScanResultsEventGridTopicResourceID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.EventGrid/topics/sampletopic"),
	// 		},
	// 		OverrideSubscriptionLevelSettings: to.Ptr(true),
	// 		SensitiveDataDiscovery: &armsecurity.SensitiveDataDiscoveryProperties{
	// 			IsEnabled: to.Ptr(false),
	// 		},
	// 	},
	// }
}
