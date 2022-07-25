const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a private link resource of a IoT Central Application.
 *
 * @summary Get a private link resource of a IoT Central Application.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_Get.json
 */
async function privateLinksGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const resourceName = "myIoTCentralApp";
  const groupId = "iotApp";
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const result = await client.privateLinks.get(resourceGroupName, resourceName, groupId);
  console.log(result);
}

privateLinksGet().catch(console.error);
