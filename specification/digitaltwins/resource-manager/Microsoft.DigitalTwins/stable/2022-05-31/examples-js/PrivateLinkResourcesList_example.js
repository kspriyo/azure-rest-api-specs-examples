const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private link resources for given Digital Twin.
 *
 * @summary List private link resources for given Digital Twin.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/PrivateLinkResourcesList_example.json
 */
async function listPrivateLinkResourcesForGivenDigitalTwin() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.list(resourceGroupName, resourceName);
  console.log(result);
}

listPrivateLinkResourcesForGivenDigitalTwin().catch(console.error);
