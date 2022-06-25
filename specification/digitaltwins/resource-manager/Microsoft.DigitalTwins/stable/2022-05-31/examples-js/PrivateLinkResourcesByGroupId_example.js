const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified private link resource for the given Digital Twin.
 *
 * @summary Get the specified private link resource for the given Digital Twin.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/PrivateLinkResourcesByGroupId_example.json
 */
async function getTheSpecifiedPrivateLinkResourceForTheGivenDigitalTwin() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const resourceId = "subResource";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, resourceName, resourceId);
  console.log(result);
}

getTheSpecifiedPrivateLinkResourceForTheGivenDigitalTwin().catch(console.error);
