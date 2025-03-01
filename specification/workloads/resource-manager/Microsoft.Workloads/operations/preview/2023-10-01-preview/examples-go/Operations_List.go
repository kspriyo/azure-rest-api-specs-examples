package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/operations/preview/2023-10-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armworkloadssapvirtualinstance.OperationListResult{
		// 	Value: []*armworkloadssapvirtualinstance.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor."),
		// 				Operation: to.Ptr("monitors_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of SAP monitors in the specified resource group."),
		// 				Operation: to.Ptr("monitors_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("monitors_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Creates a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("monitors_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SAP monitor with the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("monitors_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Patches the Tags field of a SAP monitor for the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("monitors_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of provider instances in the specified SAP monitor. The operations returns various properties of each provider instances."),
		// 				Operation: to.Ptr("ProviderInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Creates a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Deletes a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Creates a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SAP Landscape monitor configuration with the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/Write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Set SapDiscoverySites"),
		// 				Operation: to.Ptr("Creates or updates the SapDiscoverySites"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapDiscoverySites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/Delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Delete SapDiscoverySites"),
		// 				Operation: to.Ptr("Deletes the SapDiscoverySites"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapDiscoverySites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/Read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Read SapDiscoverySites"),
		// 				Operation: to.Ptr("Reads the SapDiscoverySites"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapDiscoverySites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/Write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Set SapInstances"),
		// 				Operation: to.Ptr("Creates or updates the SapInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/Delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Delete SapInstances"),
		// 				Operation: to.Ptr("Deletes the SapInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/Read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Read SapInstances"),
		// 				Operation: to.Ptr("Reads the SapInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("SapInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/serverInstances/Write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Set serverInstances"),
		// 				Operation: to.Ptr("Creates or updates the serverInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("serverInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/serverInstances/Delete"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Delete serverInstances"),
		// 				Operation: to.Ptr("Deletes the serverInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("serverInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/SapDiscoverySites/SapInstances/serverInstances/Read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Read serverInstances"),
		// 				Operation: to.Ptr("Reads the serverInstances"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("serverInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/Write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Set connectors"),
		// 				Operation: to.Ptr("Creates or updates the connectors"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/Read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Read connectors"),
		// 				Operation: to.Ptr("Reads the connectors"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/Write"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Set acssBackups"),
		// 				Operation: to.Ptr("Creates or updates the acssBackups"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/Read"),
		// 			Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 				Description: to.Ptr("Read acssBackups"),
		// 				Operation: to.Ptr("Reads the acssBackups"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
