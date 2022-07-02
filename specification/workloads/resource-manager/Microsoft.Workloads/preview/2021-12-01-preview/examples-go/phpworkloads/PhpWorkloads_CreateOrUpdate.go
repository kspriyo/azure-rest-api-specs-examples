package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_CreateOrUpdate.json
func ExamplePhpWorkloadsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewPhpWorkloadsClient("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"test-rg",
		"wp39",
		armworkloads.PhpWorkloadResource{
			Location: to.Ptr("eastus2"),
			Tags:     map[string]*string{},
			Kind:     to.Ptr(armworkloads.WorkloadKindWordPress),
			Properties: &armworkloads.PhpWorkloadResourceProperties{
				AdminUserProfile: &armworkloads.UserProfile{
					SSHPublicKey: to.Ptr("===SSH=PUBLIC=KEY==="),
					UserName:     to.Ptr("wpadmin"),
				},
				AppLocation: to.Ptr("eastus"),
				BackupProfile: &armworkloads.BackupProfile{
					BackupEnabled: to.Ptr(armworkloads.EnableBackupDisabled),
				},
				CacheProfile: &armworkloads.CacheProfile{
					Name:     to.Ptr("wp-cache"),
					Capacity: to.Ptr[int64](0),
					Family:   to.Ptr(armworkloads.RedisCacheFamilyC),
					SKUName:  to.Ptr("Basic"),
				},
				ControllerProfile: &armworkloads.NodeProfile{
					Name: to.Ptr("contoller-vm"),
					DataDisks: []*armworkloads.DiskInfo{
						{
							SizeInGB:    to.Ptr[int64](100),
							StorageType: to.Ptr(armworkloads.DiskStorageTypePremiumLRS),
						}},
					NodeSKU: to.Ptr("Standard_DS2_v2"),
					OSDisk: &armworkloads.DiskInfo{
						StorageType: to.Ptr(armworkloads.DiskStorageTypePremiumLRS),
					},
					OSImage: &armworkloads.OsImageProfile{
						Offer:     to.Ptr(armworkloads.OSImageOfferUbuntuServer),
						Publisher: to.Ptr(armworkloads.OSImagePublisherCanonical),
						SKU:       to.Ptr(armworkloads.OSImageSKU("18.0-LTS")),
						Version:   to.Ptr(armworkloads.OSImageVersionLatest),
					},
				},
				DatabaseProfile: &armworkloads.DatabaseProfile{
					Type:                  to.Ptr(armworkloads.DatabaseTypeMySQL),
					BackupRetentionDays:   to.Ptr[int32](7),
					HaEnabled:             to.Ptr(armworkloads.HAEnabledDisabled),
					ServerName:            to.Ptr("wp-db-server"),
					SKU:                   to.Ptr("Standard_D32s_v4"),
					SSLEnforcementEnabled: to.Ptr(armworkloads.EnableSSLEnforcementEnabled),
					StorageInGB:           to.Ptr[int64](128),
					StorageIops:           to.Ptr[int64](200),
					StorageSKU:            to.Ptr("Premium_LRS"),
					Tier:                  to.Ptr(armworkloads.DatabaseTierGeneralPurpose),
					Version:               to.Ptr("5.7"),
				},
				FileshareProfile: &armworkloads.FileshareProfile{
					ShareSizeInGB: to.Ptr[int64](100),
					ShareType:     to.Ptr(armworkloads.FileShareTypeAzureFiles),
					StorageType:   to.Ptr(armworkloads.FileShareStorageTypePremiumLRS),
				},
				ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
					Name: to.Ptr("php-mrg-wp39"),
				},
				NetworkProfile: &armworkloads.NetworkProfile{
					AzureFrontDoorEnabled: to.Ptr(armworkloads.AzureFrontDoorEnabledEnabled),
					LoadBalancerSKU:       to.Ptr("Standard"),
					LoadBalancerType:      to.Ptr(armworkloads.LoadBalancerTypeLoadBalancer),
				},
				PhpProfile: &armworkloads.PhpProfile{
					Version: to.Ptr(armworkloads.PHPVersionSeven3),
				},
				SearchProfile: &armworkloads.SearchProfile{
					NodeSKU: to.Ptr("Standard_DS2_v2"),
					OSDisk: &armworkloads.DiskInfo{
						StorageType: to.Ptr(armworkloads.DiskStorageTypePremiumLRS),
					},
					OSImage: &armworkloads.OsImageProfile{
						Offer:     to.Ptr(armworkloads.OSImageOfferUbuntuServer),
						Publisher: to.Ptr(armworkloads.OSImagePublisherCanonical),
						SKU:       to.Ptr(armworkloads.OSImageSKU("18.0-LTS")),
						Version:   to.Ptr(armworkloads.OSImageVersionLatest),
					},
					SearchType: to.Ptr(armworkloads.SearchTypeElastic),
				},
				SiteProfile: &armworkloads.SiteProfile{
					DomainName: to.Ptr("www.example.com"),
				},
				WebNodesProfile: &armworkloads.VmssNodesProfile{
					Name:    to.Ptr("web-server"),
					NodeSKU: to.Ptr("Standard_DS2_v2"),
					OSDisk: &armworkloads.DiskInfo{
						StorageType: to.Ptr(armworkloads.DiskStorageTypePremiumLRS),
					},
					OSImage: &armworkloads.OsImageProfile{
						Offer:     to.Ptr(armworkloads.OSImageOfferUbuntuServer),
						Publisher: to.Ptr(armworkloads.OSImagePublisherCanonical),
						SKU:       to.Ptr(armworkloads.OSImageSKU("18.0-LTS")),
						Version:   to.Ptr(armworkloads.OSImageVersionLatest),
					},
					AutoScaleMaxCount: to.Ptr[int32](1),
					AutoScaleMinCount: to.Ptr[int32](1),
				},
			},
			SKU: &armworkloads.SKU{
				Name: to.Ptr("Large"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
