package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscription_example.json
func ExampleTopologyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopologyClient().NewListPager(nil)
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
		// page.TopologyList = armsecurity.TopologyList{
		// 	Value: []*armsecurity.TopologyResource{
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("vnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/vnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.575Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Children: []*armsecurity.TopologySingleResourceChild{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						}},
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](0),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("subnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/subnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.575Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						Parents: []*armsecurity.TopologySingleResourceParent{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						}},
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](5),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
