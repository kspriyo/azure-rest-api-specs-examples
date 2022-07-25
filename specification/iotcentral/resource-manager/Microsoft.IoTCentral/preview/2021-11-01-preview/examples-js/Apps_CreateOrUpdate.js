const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the metadata of an IoT Central application. The usual pattern to modify a property is to retrieve the IoT Central application metadata and security metadata, and then combine them with the modified values in a new body to update the IoT Central application.
 *
 * @summary Create or update the metadata of an IoT Central application. The usual pattern to modify a property is to retrieve the IoT Central application metadata and security metadata, and then combine them with the modified values in a new body to update the IoT Central application.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CreateOrUpdate.json
 */
async function appsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const resourceName = "myIoTCentralApp";
  const app = {
    displayName: "My IoT Central App",
    identity: { type: "SystemAssigned" },
    location: "westus",
    sku: { name: "ST2" },
    subdomain: "my-iot-central-app",
    template: "iotc-pnp-preview@1.0.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const result = await client.apps.beginCreateOrUpdateAndWait(resourceGroupName, resourceName, app);
  console.log(result);
}

appsCreateOrUpdate().catch(console.error);
