package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/GetExternalSecuritySolutionsSubscriptionLocation_example.json
func ExampleExternalSecuritySolutionsClient_NewListByHomeRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalSecuritySolutionsClient().NewListByHomeRegionPager("centralus", nil)
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
		// page.ExternalSecuritySolutionList = armsecurity.ExternalSecuritySolutionList{
		// 	Value: []armsecurity.ExternalSecuritySolutionClassification{
		// 		&armsecurity.AADExternalSecuritySolution{
		// 			Location: to.Ptr("eastus"),
		// 			Name: to.Ptr("aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/defaultresourcegroup-eus/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindAAD),
		// 			Properties: &armsecurity.AADSolutionProperties{
		// 				DeviceType: to.Ptr("Azure Active Directory Identity Protection"),
		// 				DeviceVendor: to.Ptr("Microsoft"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/defaultresourcegroup-eus/providers/Microsoft.OperationalInsights/workspaces/defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
		// 				},
		// 				AdditionalProperties: map[string]any{
		// 					"connectivityState": "Discovered",
		// 				},
		// 			},
		// 		},
		// 		&armsecurity.AADExternalSecuritySolution{
		// 			Location: to.Ptr("westeurope"),
		// 			Name: to.Ptr("aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-weu"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/defaultresourcegroup-weu/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-weu"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindAAD),
		// 			Properties: &armsecurity.AADSolutionProperties{
		// 				DeviceType: to.Ptr("Azure Active Directory Identity Protection"),
		// 				DeviceVendor: to.Ptr("Microsoft"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/defaultresourcegroup-weu/providers/Microsoft.OperationalInsights/workspaces/defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-weu"),
		// 				},
		// 				AdditionalProperties: map[string]any{
		// 					"connectivityState": "Discovered",
		// 				},
		// 			},
		// 		},
		// 		&armsecurity.CefExternalSecuritySolution{
		// 			Location: to.Ptr("westcentralus"),
		// 			Name: to.Ptr("cef_omsprd_barracudanetworks_waf_barracuda"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/unificationprod/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/cef_omsprd_barracudanetworks_waf_barracuda"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindCEF),
		// 			Properties: &armsecurity.CefSolutionProperties{
		// 				DeviceType: to.Ptr("WAF"),
		// 				DeviceVendor: to.Ptr("barracudanetworks"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/unificationprod/providers/Microsoft.OperationalInsights/workspaces/omsprd"),
		// 				},
		// 				Hostname: to.Ptr("barracuda"),
		// 				LastEventReceived: to.Ptr("2018-05-09T10:30:11.523Z"),
		// 			},
		// 		},
		// 		&armsecurity.CefExternalSecuritySolution{
		// 			Location: to.Ptr("westcentralus"),
		// 			Name: to.Ptr("cef_omsprd_virtualhoneypot_Microsoft_demovm20"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/unificationprod/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/cef_omsprd_virtualhoneypot_Microsoft_demovm20"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindCEF),
		// 			Properties: &armsecurity.CefSolutionProperties{
		// 				DeviceType: to.Ptr("Microsoft"),
		// 				DeviceVendor: to.Ptr("virtualHoneypot"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/unificationprod/providers/Microsoft.OperationalInsights/workspaces/omsprd"),
		// 				},
		// 				Hostname: to.Ptr("demovm20"),
		// 				LastEventReceived: to.Ptr("2018-05-08T15:42:22.57Z"),
		// 			},
		// 		},
		// 		&armsecurity.CefExternalSecuritySolution{
		// 			Location: to.Ptr("westcentralus"),
		// 			Name: to.Ptr("cef_omsprd_virtualhoneypot_Microsoft_demovm10"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/unificationprod/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/cef_omsprd_virtualhoneypot_Microsoft_demovm10"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindCEF),
		// 			Properties: &armsecurity.CefSolutionProperties{
		// 				DeviceType: to.Ptr("Microsoft"),
		// 				DeviceVendor: to.Ptr("virtualHoneypot"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/unificationprod/providers/Microsoft.OperationalInsights/workspaces/omsprd"),
		// 				},
		// 				Hostname: to.Ptr("demovm10"),
		// 				LastEventReceived: to.Ptr("2018-05-08T10:38:53.423Z"),
		// 			},
		// 		},
		// 		&armsecurity.AADExternalSecuritySolution{
		// 			Location: to.Ptr("westcentralus"),
		// 			Name: to.Ptr("aad_omsprd"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/unificationprod/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/aad_omsprd"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindAAD),
		// 			Properties: &armsecurity.AADSolutionProperties{
		// 				DeviceType: to.Ptr("Azure Active Directory Identity Protection"),
		// 				DeviceVendor: to.Ptr("Microsoft"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/unificationprod/providers/Microsoft.OperationalInsights/workspaces/omsprd"),
		// 				},
		// 				AdditionalProperties: map[string]any{
		// 					"connectivityState": "Discovered",
		// 				},
		// 			},
		// 		},
		// 		&armsecurity.AADExternalSecuritySolution{
		// 			Location: to.Ptr("japaneast"),
		// 			Name: to.Ptr("aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-ejp"),
		// 			Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/defaultresourcegroup-ejp/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-ejp"),
		// 			Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindAAD),
		// 			Properties: &armsecurity.AADSolutionProperties{
		// 				DeviceType: to.Ptr("Azure Active Directory Identity Protection"),
		// 				DeviceVendor: to.Ptr("Microsoft"),
		// 				Workspace: &armsecurity.ConnectedWorkspace{
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/defaultresourcegroup-ejp/providers/Microsoft.OperationalInsights/workspaces/defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-ejp"),
		// 				},
		// 				AdditionalProperties: map[string]any{
		// 					"connectivityState": "Discovered",
		// 				},
		// 			},
		// 	}},
		// }
	}
}
