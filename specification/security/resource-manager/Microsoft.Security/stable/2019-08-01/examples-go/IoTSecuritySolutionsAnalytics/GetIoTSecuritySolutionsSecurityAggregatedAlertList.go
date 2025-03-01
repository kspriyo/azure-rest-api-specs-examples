package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlertList.json
func ExampleIotSecuritySolutionsAnalyticsAggregatedAlertClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotSecuritySolutionsAnalyticsAggregatedAlertClient().NewListPager("MyGroup", "default", &armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions{Top: nil})
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
		// page.IoTSecurityAggregatedAlertList = armsecurity.IoTSecurityAggregatedAlertList{
		// 	Value: []*armsecurity.IoTSecurityAggregatedAlert{
		// 		{
		// 			Name: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Fail/2019-02-02"),
		// 			Type: to.Ptr("Microsoft.Security/iotSecuritySolutions/analyticsModels/aggregatedAlerts"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Fail/2019-02-02"),
		// 			Properties: &armsecurity.IoTSecurityAggregatedAlertProperties{
		// 				Description: to.Ptr("Multiple unsuccsseful login attempts identified. A Bruteforce attack on the device failed."),
		// 				ActionTaken: to.Ptr("Detected"),
		// 				AggregatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-02-02"); return t}()),
		// 				AlertDisplayName: to.Ptr("Failed Bruteforce"),
		// 				AlertType: to.Ptr("IoT_Bruteforce_Fail"),
		// 				Count: to.Ptr[int64](50),
		// 				EffectedResourceType: to.Ptr("IoT Device"),
		// 				LogAnalyticsQuery: to.Ptr("SecurityAlert | where tolower(ResourceId) == tolower('/subscriptions/b77ec8a9-04ed-48d2-a87a-e5887b978ba6/resourceGroups/IoT-Solution-DemoEnv/providers/Microsoft.Devices/IotHubs/rtogm-hub') and tolower(AlertName) == tolower('Custom Alert - number of device to cloud messages in MQTT protocol is not in the allowed range') | extend DeviceId=parse_json(ExtendedProperties)['DeviceId'] | project DeviceId, TimeGenerated, DisplayName, AlertSeverity, Description, RemediationSteps, ExtendedProperties"),
		// 				RemediationSteps: to.Ptr(""),
		// 				ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
		// 				SystemSource: to.Ptr("Devices"),
		// 				TopDevicesList: []*armsecurity.IoTSecurityAggregatedAlertPropertiesTopDevicesListItem{
		// 					{
		// 						AlertsCount: to.Ptr[int64](45),
		// 						DeviceID: to.Ptr("testDevice1"),
		// 						LastOccurrence: to.Ptr("10:42"),
		// 					},
		// 					{
		// 						AlertsCount: to.Ptr[int64](30),
		// 						DeviceID: to.Ptr("testDevice2"),
		// 						LastOccurrence: to.Ptr("15:42"),
		// 				}},
		// 				VendorName: to.Ptr("Microsoft"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Success/2019-02-02"),
		// 			Type: to.Ptr("Microsoft.Security/iotSecuritySolutions/analyticsModels/aggregatedAlerts"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Success/2019-02-02"),
		// 			Properties: &armsecurity.IoTSecurityAggregatedAlertProperties{
		// 				Description: to.Ptr("Multiple unsuccsseful login attempts identified followed by a succssful login. A Bruteforce attack on the device was Successfule"),
		// 				ActionTaken: to.Ptr("Detected"),
		// 				AggregatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-02-02"); return t}()),
		// 				AlertDisplayName: to.Ptr("Successful Bruteforce"),
		// 				AlertType: to.Ptr("IoT_Bruteforce_Success"),
		// 				Count: to.Ptr[int64](600000),
		// 				EffectedResourceType: to.Ptr("IoT Device"),
		// 				LogAnalyticsQuery: to.Ptr("SecurityAlert | where tolower(ResourceId) == tolower('/subscriptions/b77ec8a9-04ed-48d2-a87a-e5887b978ba6/resourceGroups/IoT-Solution-DemoEnv/providers/Microsoft.Devices/IotHubs/rtogm-hub') and tolower(AlertName) == tolower('Custom Alert - number of device to cloud messages in MQTT protocol is not in the allowed range') | extend DeviceId=parse_json(ExtendedProperties)['DeviceId'] | project DeviceId, TimeGenerated, DisplayName, AlertSeverity, Description, RemediationSteps, ExtendedProperties"),
		// 				RemediationSteps: to.Ptr(""),
		// 				ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
		// 				SystemSource: to.Ptr("Devices"),
		// 				TopDevicesList: []*armsecurity.IoTSecurityAggregatedAlertPropertiesTopDevicesListItem{
		// 					{
		// 						AlertsCount: to.Ptr[int64](12321),
		// 						DeviceID: to.Ptr("testDevice1"),
		// 						LastOccurrence: to.Ptr("10:42"),
		// 					},
		// 					{
		// 						AlertsCount: to.Ptr[int64](455),
		// 						DeviceID: to.Ptr("testDevice2"),
		// 						LastOccurrence: to.Ptr("15:42"),
		// 				}},
		// 				VendorName: to.Ptr("Microsoft"),
		// 			},
		// 	}},
		// }
	}
}
