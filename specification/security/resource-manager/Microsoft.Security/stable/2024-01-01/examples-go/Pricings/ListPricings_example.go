package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/ListPricings_example.json
func ExamplePricingsClient_List_getPricingsOnSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().List(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", &armsecurity.PricingsClientListOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PricingList = armsecurity.PricingList{
	// 	Value: []*armsecurity.Pricing{
	// 		{
	// 			Name: to.Ptr("VirtualMachines"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/VirtualMachines"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusPartiallyCovered),
	// 				SubPlan: to.Ptr("P2"),
	// 				Extensions: []*armsecurity.Extension{
	// 					{
	// 						Name: to.Ptr("AgentlessVmScanning"),
	// 						AdditionalExtensionProperties: map[string]any{
	// 							"ExclusionTags": "[{\"Key\":\"TestKey1\",\"Value\":\"TestValue1\"},{\"Key\":\"TestKey2\",\"Value\":\"TestValue2\"}]",
	// 						},
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 					},
	// 					{
	// 						Name: to.Ptr("MdeDesignatedSubscription"),
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				}},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("SqlServers"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/SqlServers"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("AppServices"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/AppServices"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 				ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusNotCovered),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("StorageAccounts"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/StorageAccounts"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 				SubPlan: to.Ptr("DefenderForStorageV2"),
	// 				Extensions: []*armsecurity.Extension{
	// 					{
	// 						Name: to.Ptr("OnUploadMalwareScanning"),
	// 						AdditionalExtensionProperties: map[string]any{
	// 							"capGBPerMonthPerStorageAccount": float64(10),
	// 						},
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 					},
	// 					{
	// 						Name: to.Ptr("SensitiveDataDiscovery"),
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				}},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("SqlServerVirtualMachines"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/SqlServerVirtualMachines"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("KubernetesService"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/KubernetesService"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				Deprecated: to.Ptr(true),
	// 				Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 				ReplacedBy: []*string{
	// 					to.Ptr("Containers")},
	// 					ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusNotCovered),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("ContainerRegistry"),
	// 				Type: to.Ptr("Microsoft.Security/pricings"),
	// 				ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/ContainerRegistry"),
	// 				Properties: &armsecurity.PricingProperties{
	// 					Deprecated: to.Ptr(true),
	// 					Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 					FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 					PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 					ReplacedBy: []*string{
	// 						to.Ptr("Containers")},
	// 						ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusNotCovered),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("KeyVaults"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/KeyVaults"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 						ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 						SubPlan: to.Ptr("PerKeyVault"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Dns"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Dns"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						Deprecated: to.Ptr(true),
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 						ReplacedBy: []*string{
	// 							to.Ptr("VirtualMachines")},
	// 							ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("Arm"),
	// 						Type: to.Ptr("Microsoft.Security/pricings"),
	// 						ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Arm"),
	// 						Properties: &armsecurity.PricingProperties{
	// 							EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 							Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 							FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 							PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 							ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 							SubPlan: to.Ptr("PerSubscription"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("OpenSourceRelationalDatabases"),
	// 						Type: to.Ptr("Microsoft.Security/pricings"),
	// 						ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/OpenSourceRelationalDatabases"),
	// 						Properties: &armsecurity.PricingProperties{
	// 							Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 							FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 							PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 							ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("Containers"),
	// 						Type: to.Ptr("Microsoft.Security/pricings"),
	// 						ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Containers"),
	// 						Properties: &armsecurity.PricingProperties{
	// 							EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 							Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 							FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 							PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 							ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 							Extensions: []*armsecurity.Extension{
	// 								{
	// 									Name: to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CloudPosture"),
	// 						Type: to.Ptr("Microsoft.Security/pricings"),
	// 						ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/CloudPosture"),
	// 						Properties: &armsecurity.PricingProperties{
	// 							EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 							Enforce: to.Ptr(armsecurity.EnforceFalse),
	// 							FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 							PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 							ResourcesCoverageStatus: to.Ptr(armsecurity.ResourcesCoverageStatusFullyCovered),
	// 							Extensions: []*armsecurity.Extension{
	// 								{
	// 									Name: to.Ptr("AgentlessVmScanning"),
	// 									AdditionalExtensionProperties: map[string]any{
	// 										"ExclusionTags": "[]",
	// 									},
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 								},
	// 								{
	// 									Name: to.Ptr("AgentlessDiscoveryForKubernetes"),
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 								},
	// 								{
	// 									Name: to.Ptr("SensitiveDataDiscovery"),
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 								},
	// 								{
	// 									Name: to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 								},
	// 								{
	// 									Name: to.Ptr("EntraPermissionsManagement"),
	// 									IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("Api"),
	// 						Type: to.Ptr("Microsoft.Security/pricings"),
	// 						ID: to.Ptr("subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Api"),
	// 						Properties: &armsecurity.PricingProperties{
	// 							FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 							PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 							SubPlan: to.Ptr("P1"),
	// 						},
	// 				}},
	// 			}
}
