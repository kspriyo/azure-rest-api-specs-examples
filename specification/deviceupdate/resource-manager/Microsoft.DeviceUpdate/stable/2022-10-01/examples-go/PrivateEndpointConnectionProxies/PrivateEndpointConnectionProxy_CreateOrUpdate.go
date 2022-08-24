package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeviceupdate.NewPrivateEndpointConnectionProxiesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointConnectionProxy{
		RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
			ID:                      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Location:                to.Ptr("westus2"),
			ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
				{
					Name: to.Ptr("{privateEndpointConnectionProxyId}"),
					GroupIDs: []*string{
						to.Ptr("DeviceUpdate")},
					RequestMessage: to.Ptr("Please approve my connection, thanks."),
				}},
			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
				{
					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
					ID:                           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
