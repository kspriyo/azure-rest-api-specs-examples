package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/SensitivitySettings/GetSensitivitySettingsList_example.json
func ExampleSensitivitySettingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSensitivitySettingsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetSensitivitySettingsListResponse = armsecurity.GetSensitivitySettingsListResponse{
	// 	Value: []*armsecurity.GetSensitivitySettingsResponse{
	// 		{
	// 			Name: to.Ptr("current"),
	// 			Type: to.Ptr("Microsoft.Security/sensitivitySettings"),
	// 			ID: to.Ptr("/providers/Microsoft.Security/sensitivitySettings"),
	// 			Properties: &armsecurity.GetSensitivitySettingsResponseProperties{
	// 				MipInformation: &armsecurity.GetSensitivitySettingsResponsePropertiesMipInformation{
	// 					BuiltInInfoTypes: []*armsecurity.BuiltInInfoType{
	// 						{
	// 							Name: to.Ptr("Http Authorization Header"),
	// 							Type: to.Ptr("Credentials"),
	// 							ID: to.Ptr("4d0d3eb6-619f-4c8c-810c-c16150c95278"),
	// 						},
	// 						{
	// 							Name: to.Ptr("User Login Credentials"),
	// 							Type: to.Ptr("Credentials"),
	// 							ID: to.Ptr("a98fde82-45b6-4b2c-afd0-ad579cd9f826"),
	// 						},
	// 						{
	// 							Name: to.Ptr("Credit card number"),
	// 							Type: to.Ptr("Finance"),
	// 							ID: to.Ptr("50842eb7-edc8-4019-85dd-5a5c1f2bb085"),
	// 						},
	// 						{
	// 							Name: to.Ptr("EU debit card number"),
	// 							Type: to.Ptr("Finance"),
	// 							ID: to.Ptr("0e9b3178-9678-47dd-a509-37222ca96b42"),
	// 					}},
	// 					CustomInfoTypes: []*armsecurity.InfoType{
	// 						{
	// 							Name: to.Ptr("User created custom info type 1"),
	// 							Description: to.Ptr("Custom info type description"),
	// 							ID: to.Ptr("c5f9b9a1-2b9a-4a3a-8c5b-1f3d1d9d9c9b"),
	// 						},
	// 						{
	// 							Name: to.Ptr("User created custom info type 2"),
	// 							Description: to.Ptr("Custom info type description"),
	// 							ID: to.Ptr("a4fee2b6-5618-404b-a5e7-aa377cd67543"),
	// 						},
	// 						{
	// 							Name: to.Ptr("User created custom info type 3"),
	// 							Description: to.Ptr("Custom info type description"),
	// 							ID: to.Ptr("a355f11e-f87d-4f48-8490-ecf0873325b5"),
	// 					}},
	// 					Labels: []*armsecurity.Label{
	// 						{
	// 							Name: to.Ptr("Public"),
	// 							ID: to.Ptr("fdfb5435-124d-4651-a889-a4210fec6a77"),
	// 							Order: to.Ptr[float32](0),
	// 						},
	// 						{
	// 							Name: to.Ptr("Confidential"),
	// 							ID: to.Ptr("4c5447ec-f7f3-4345-a160-6a5850f2bf0c"),
	// 							Order: to.Ptr[float32](1),
	// 						},
	// 						{
	// 							Name: to.Ptr("Highly Confidential"),
	// 							ID: to.Ptr("f38ac75c-f42a-4c89-aa37-9c4e74101414"),
	// 							Order: to.Ptr[float32](2),
	// 					}},
	// 					MipIntegrationStatus: to.Ptr(armsecurity.MipIntegrationStatusOk),
	// 				},
	// 				SensitiveInfoTypesIDs: []*string{
	// 					to.Ptr("a98fde82-45b6-4b2c-afd0-ad579cd9f826"),
	// 					to.Ptr("4d0d3eb6-619f-4c8c-810c-c16150c95278"),
	// 					to.Ptr("0e9b3178-9678-47dd-a509-37222ca96b42"),
	// 					to.Ptr("c5f9b9a1-2b9a-4a3a-8c5b-1f3d1d9d9c9b")},
	// 					SensitivityThresholdLabelOrder: to.Ptr[float32](1),
	// 				},
	// 		}},
	// 	}
}
