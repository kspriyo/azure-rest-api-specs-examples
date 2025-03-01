package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetIoTDeviceEntityById.json
func ExampleEntitiesClient_Get_getAnIoTDeviceEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.IoTDeviceEntity{
	// 		Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindIoTDevice),
	// 		Properties: &armsecurityinsights.IoTDeviceEntityProperties{
	// 			FriendlyName: to.Ptr("device1"),
	// 			DeviceID: to.Ptr("device1"),
	// 			DeviceName: to.Ptr("device1"),
	// 			DeviceType: to.Ptr("Industrial"),
	// 			FirmwareVersion: to.Ptr("20.11"),
	// 			Importance: to.Ptr(armsecurityinsights.DeviceImportanceNormal),
	// 			IotHubEntityID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/8b2d9401-f953-e89d-2583-be9b4975870c"),
	// 			IsAuthorized: to.Ptr(true),
	// 			IsProgramming: to.Ptr(false),
	// 			IsScanner: to.Ptr(false),
	// 			Model: to.Ptr("demo-model"),
	// 			NicEntityIDs: []*string{
	// 				to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/6ee379bd-ace8-44cf-ab10-ee669a1b71e2")},
	// 				OperatingSystem: to.Ptr("Windows"),
	// 				Protocols: []*string{
	// 					to.Ptr("CIP"),
	// 					to.Ptr("EtherNet/IP")},
	// 					PurdueLayer: to.Ptr("ProcessControl"),
	// 					Sensor: to.Ptr("demo-sensor"),
	// 					Site: to.Ptr("demo-site"),
	// 					Vendor: to.Ptr("demo-vendor"),
	// 					Zone: to.Ptr("zone"),
	// 				},
	// 			},
	// 			                        }
}
