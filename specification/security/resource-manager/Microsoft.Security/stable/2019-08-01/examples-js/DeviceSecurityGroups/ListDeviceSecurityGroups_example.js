const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method get the list of device security groups for the specified IoT Hub resource.
 *
 * @summary Use this method get the list of device security groups for the specified IoT Hub resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/ListDeviceSecurityGroups_example.json
 */
async function listAllDeviceSecurityGroupsForTheSpecifiedIoTHubResource() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deviceSecurityGroups.list(resourceId)) {
    resArray.push(item);
  }
  console.log(resArray);
}
