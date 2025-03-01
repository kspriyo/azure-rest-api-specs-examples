package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/GetAllowedConnectionsSubscription_example.json
func ExampleAllowedConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAllowedConnectionsClient().NewListPager(nil)
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
		// page.AllowedConnectionsList = armsecurity.AllowedConnectionsList{
		// 	Value: []*armsecurity.AllowedConnectionsResource{
		// 		{
		// 			Location: to.Ptr("centralus"),
		// 			Name: to.Ptr("Internal"),
		// 			Type: to.Ptr("Microsoft.Security/locations/allowedConnections"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Security/locations/centralus/allowedConnections/Internal"),
		// 			Properties: &armsecurity.AllowedConnectionsResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-06T14:55:32.351Z"); return t}()),
		// 				ConnectableResources: []*armsecurity.ConnectableResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine1"),
		// 						InboundConnectedResources: []*armsecurity.ConnectedResource{
		// 							{
		// 								ConnectedResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine2"),
		// 								TCPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 								UDPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 						}},
		// 						OutboundConnectedResources: []*armsecurity.ConnectedResource{
		// 							{
		// 								ConnectedResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine2"),
		// 								TCPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 								UDPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 						}},
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine2"),
		// 						InboundConnectedResources: []*armsecurity.ConnectedResource{
		// 							{
		// 								ConnectedResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine1"),
		// 								TCPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 								UDPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 						}},
		// 						OutboundConnectedResources: []*armsecurity.ConnectedResource{
		// 							{
		// 								ConnectedResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine1"),
		// 								TCPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 								UDPPorts: to.Ptr("[0-21,23-3388,3390-5984,5987-65535]"),
		// 						}},
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154baf/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtaulMachine3"),
		// 						InboundConnectedResources: []*armsecurity.ConnectedResource{
		// 						},
		// 						OutboundConnectedResources: []*armsecurity.ConnectedResource{
		// 						},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
