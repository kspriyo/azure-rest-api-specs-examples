package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/ListDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeviceSecurityGroupsClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub", nil)
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
		// page.DeviceSecurityGroupList = armsecurity.DeviceSecurityGroupList{
		// 	Value: []*armsecurity.DeviceSecurityGroup{
		// 		{
		// 			Name: to.Ptr("samplesecuritygroup"),
		// 			Type: to.Ptr("Microsoft.Security/deviceSecurityGroups"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub/providers/Microsoft.Security/deviceSecurityGroups/samplesecuritygroup"),
		// 			Properties: &armsecurity.DeviceSecurityGroupProperties{
		// 				AllowlistRules: []armsecurity.AllowlistCustomAlertRuleClassification{
		// 					&armsecurity.ConnectionToIPNotAllowed{
		// 						Description: to.Ptr("Get an alert when an outbound connection is created between your device and an ip that isn't allowed"),
		// 						DisplayName: to.Ptr("Outbound connection to an ip that isn't allowed"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("ConnectionToIpNotAllowed"),
		// 						ValueType: to.Ptr(armsecurity.ValueTypeIPCidr),
		// 						AllowlistValues: []*string{
		// 						},
		// 					},
		// 					&armsecurity.LocalUserNotAllowed{
		// 						Description: to.Ptr("Get an alert when a local user that isn't allowed logins to the device"),
		// 						DisplayName: to.Ptr("Login by a local user that isn't allowed"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("LocalUserNotAllowed"),
		// 						ValueType: to.Ptr(armsecurity.ValueTypeString),
		// 						AllowlistValues: []*string{
		// 						},
		// 					},
		// 					&armsecurity.ProcessNotAllowed{
		// 						Description: to.Ptr("Get an alert when a process that isn't allowed is executed"),
		// 						DisplayName: to.Ptr("Execution of a process that isn't allowed"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("ProcessNotAllowed"),
		// 						ValueType: to.Ptr(armsecurity.ValueTypeString),
		// 						AllowlistValues: []*string{
		// 						},
		// 				}},
		// 				DenylistRules: []*armsecurity.DenylistCustomAlertRule{
		// 				},
		// 				ThresholdRules: []armsecurity.ThresholdCustomAlertRuleClassification{
		// 				},
		// 				TimeWindowRules: []armsecurity.TimeWindowCustomAlertRuleClassification{
		// 					&armsecurity.ActiveConnectionsNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of active connections of a device in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of active connections is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("ActiveConnectionsNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.AmqpC2DMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (AMQP protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of cloud to device messages (AMQP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("AmqpC2DMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.MqttC2DMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (MQTT protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of cloud to device messages (MQTT protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("MqttC2DMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.HTTPC2DMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (HTTP protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of cloud to device messages (HTTP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("HttpC2DMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.AmqpC2DRejectedMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (AMQP protocol) that were rejected by the device in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of rejected cloud to device messages (AMQP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("AmqpC2DRejectedMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.MqttC2DRejectedMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (MQTT protocol) that were rejected by the device in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of rejected cloud to device messages (MQTT protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("MqttC2DRejectedMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.HTTPC2DRejectedMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of cloud to device messages (HTTP protocol) that were rejected by the device in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of rejected cloud to device messages (HTTP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("HttpC2DRejectedMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.AmqpD2CMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of device to cloud messages (AMQP protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of device to cloud messages (AMQP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("AmqpD2CMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.MqttD2CMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of device to cloud messages (MQTT protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of device to cloud messages (MQTT protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("MqttD2CMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.HTTPD2CMessagesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of device to cloud messages (HTTP protocol) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of device to cloud messages (HTTP protocol) is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("HttpD2CMessagesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.DirectMethodInvokesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of direct method invokes in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of direct method invokes is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("DirectMethodInvokesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.FailedLocalLoginsNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of failed local logins on the device in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of failed local logins is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("FailedLocalLoginsNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.FileUploadsNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of file uploads from the device to the cloud in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of file uploads is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("FileUploadsNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.QueuePurgesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of device queue purges in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of device queue purges is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("QueuePurgesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.TwinUpdatesNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number of twin updates (by the device or the service) in the time window is not in the allowed range"),
		// 						DisplayName: to.Ptr("Number of twin updates is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("TwinUpdatesNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 					},
		// 					&armsecurity.UnauthorizedOperationsNotInAllowedRange{
		// 						Description: to.Ptr("Get an alert when the number unauthorized operations in the time window is not in the allowed range. Unauthorized operations are operations that affect the device (or done by it) that fail because of an unauthorized error"),
		// 						DisplayName: to.Ptr("Number of unauthorized operations is not in allowed range"),
		// 						IsEnabled: to.Ptr(false),
		// 						RuleType: to.Ptr("UnauthorizedOperationsNotInAllowedRange"),
		// 						MaxThreshold: to.Ptr[int32](0),
		// 						MinThreshold: to.Ptr[int32](0),
		// 						TimeWindowSize: to.Ptr("PT15M"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
