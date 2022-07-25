const { OpenEnergyPlatformManagementServiceAPIs } = require("@azure/arm-oep");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of oep resources under the given Azure Subscription ID.
 *
 * @summary Lists a collection of oep resources under the given Azure Subscription ID.
 * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_ListBySubscriptionId.json
 */
async function oepResourceListBySubscriptionId() {
  const subscriptionId = "0000000-0000-0000-0000-000000000001";
  const credential = new DefaultAzureCredential();
  const client = new OpenEnergyPlatformManagementServiceAPIs(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.energyServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

oepResourceListBySubscriptionId().catch(console.error);
