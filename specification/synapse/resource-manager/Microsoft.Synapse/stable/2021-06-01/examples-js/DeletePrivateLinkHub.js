const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a privateLinkHub
 *
 * @summary Deletes a privateLinkHub
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeletePrivateLinkHub.json
 */
async function deleteAPrivateLinkHub() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const privateLinkHubName = "privateLinkHub1";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.privateLinkHubs.beginDeleteAndWait(
    resourceGroupName,
    privateLinkHubName
  );
  console.log(result);
}

deleteAPrivateLinkHub().catch(console.error);
