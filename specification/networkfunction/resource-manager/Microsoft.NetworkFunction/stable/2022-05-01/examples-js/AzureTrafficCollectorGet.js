const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Azure Traffic Collector in a specified resource group
 *
 * @summary Gets the specified Azure Traffic Collector in a specified resource group
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorGet.json
 */
async function getTrafficCollector() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.azureTrafficCollectors.get(
    resourceGroupName,
    azureTrafficCollectorName
  );
  console.log(result);
}

getTrafficCollector().catch(console.error);
