const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a privateLinkHub
 *
 * @summary Creates or updates a privateLinkHub
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdatePrivateLinkHub.json
 */
async function createOrUpdateAPrivateLinkHub() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const privateLinkHubName = "privateLinkHub1";
  const privateLinkHubInfo = {
    location: "East US",
    tags: { key: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.privateLinkHubs.createOrUpdate(
    resourceGroupName,
    privateLinkHubName,
    privateLinkHubInfo
  );
  console.log(result);
}

createOrUpdateAPrivateLinkHub().catch(console.error);
