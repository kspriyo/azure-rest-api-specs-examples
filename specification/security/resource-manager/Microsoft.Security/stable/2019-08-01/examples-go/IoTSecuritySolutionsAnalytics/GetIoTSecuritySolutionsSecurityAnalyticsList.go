package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAnalyticsList.json
func ExampleIotSecuritySolutionAnalyticsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotSecuritySolutionAnalyticsClient().List(ctx, "MyGroup", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IoTSecuritySolutionAnalyticsModelList = armsecurity.IoTSecuritySolutionAnalyticsModelList{
	// 	Value: []*armsecurity.IoTSecuritySolutionAnalyticsModel{
	// 		{
	// 			Name: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default"),
	// 			Type: to.Ptr("Microsoft.Security/iotSecuritySolutions/analyticsModels"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default"),
	// 			Properties: &armsecurity.IoTSecuritySolutionAnalyticsModelProperties{
	// 				DevicesMetrics: []*armsecurity.IoTSecuritySolutionAnalyticsModelPropertiesDevicesMetricsItem{
	// 					{
	// 						Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-01T00:00:00.000Z"); return t}()),
	// 						DevicesMetrics: &armsecurity.IoTSeverityMetrics{
	// 							High: to.Ptr[int64](3),
	// 							Low: to.Ptr[int64](70),
	// 							Medium: to.Ptr[int64](15),
	// 						},
	// 					},
	// 					{
	// 						Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-02T00:00:00.000Z"); return t}()),
	// 						DevicesMetrics: &armsecurity.IoTSeverityMetrics{
	// 							High: to.Ptr[int64](3),
	// 							Low: to.Ptr[int64](65),
	// 							Medium: to.Ptr[int64](45),
	// 						},
	// 				}},
	// 				Metrics: &armsecurity.IoTSeverityMetrics{
	// 					High: to.Ptr[int64](5),
	// 					Low: to.Ptr[int64](102),
	// 					Medium: to.Ptr[int64](200),
	// 				},
	// 				MostPrevalentDeviceAlerts: []*armsecurity.IoTSecurityDeviceAlert{
	// 					{
	// 						AlertDisplayName: to.Ptr("Custom Alert - number of device to cloud messages in AMQP protocol is not in the allowed range"),
	// 						AlertsCount: to.Ptr[int64](200),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 					},
	// 					{
	// 						AlertDisplayName: to.Ptr("Custom Alert - execution of a process that is not allowed"),
	// 						AlertsCount: to.Ptr[int64](170),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityMedium),
	// 					},
	// 					{
	// 						AlertDisplayName: to.Ptr("Successful Bruteforce"),
	// 						AlertsCount: to.Ptr[int64](150),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 				}},
	// 				MostPrevalentDeviceRecommendations: []*armsecurity.IoTSecurityDeviceRecommendation{
	// 					{
	// 						DevicesCount: to.Ptr[int64](200),
	// 						RecommendationDisplayName: to.Ptr("Install the Azure Security of Things Agent"),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 					},
	// 					{
	// 						DevicesCount: to.Ptr[int64](170),
	// 						RecommendationDisplayName: to.Ptr("High level permissions configured in Edge model twin for Edge module"),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 					},
	// 					{
	// 						DevicesCount: to.Ptr[int64](150),
	// 						RecommendationDisplayName: to.Ptr("Same Authentication Credentials used by multiple devices"),
	// 						ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityMedium),
	// 				}},
	// 				TopAlertedDevices: []*armsecurity.IoTSecurityAlertedDevice{
	// 					{
	// 						AlertsCount: to.Ptr[int64](200),
	// 						DeviceID: to.Ptr("id1"),
	// 					},
	// 					{
	// 						AlertsCount: to.Ptr[int64](170),
	// 						DeviceID: to.Ptr("id2"),
	// 					},
	// 					{
	// 						AlertsCount: to.Ptr[int64](150),
	// 						DeviceID: to.Ptr("id3"),
	// 				}},
	// 				UnhealthyDeviceCount: to.Ptr[int64](1200),
	// 			},
	// 	}},
	// }
}
