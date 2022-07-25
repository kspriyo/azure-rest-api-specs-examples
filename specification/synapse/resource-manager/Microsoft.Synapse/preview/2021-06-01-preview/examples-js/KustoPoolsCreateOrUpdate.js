const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Kusto pool.
 *
 * @summary Create or update a Kusto pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCreateOrUpdate.json
 */
async function kustoPoolsCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const workspaceName = "synapseWorkspaceName";
  const resourceGroupName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const parameters = {
    enablePurge: true,
    enableStreamingIngest: true,
    location: "westus",
    sku: { name: "Storage optimized", capacity: 2, size: "Medium" },
    workspaceUID: "11111111-2222-3333-444444444444",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.beginCreateOrUpdateAndWait(
    workspaceName,
    resourceGroupName,
    kustoPoolName,
    parameters
  );
  console.log(result);
}

kustoPoolsCreateOrUpdate().catch(console.error);
