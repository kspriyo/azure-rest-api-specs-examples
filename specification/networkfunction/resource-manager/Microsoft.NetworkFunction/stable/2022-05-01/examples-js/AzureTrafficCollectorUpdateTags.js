const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified Azure Traffic Collector tags.
 *
 * @summary Updates the specified Azure Traffic Collector tags.
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorUpdateTags.json
 */
async function updateTrafficCollectorTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const parameters = { tags: { key1: "value1", key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.azureTrafficCollectors.updateTags(
    resourceGroupName,
    azureTrafficCollectorName,
    parameters
  );
  console.log(result);
}

updateTrafficCollectorTags().catch(console.error);
