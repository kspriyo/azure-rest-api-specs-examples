const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Kusto cluster.
 *
 * @summary Create or update a Kusto cluster.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClustersCreateOrUpdate.json
 */
async function kustoClustersCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const parameters = {
    allowedIpRangeList: ["0.0.0.0/0"],
    enableAutoStop: true,
    enableDoubleEncryption: false,
    enablePurge: true,
    enableStreamingIngest: true,
    identity: { type: "SystemAssigned" },
    location: "westus",
    publicIPType: "DualStack",
    publicNetworkAccess: "Enabled",
    sku: { name: "Standard_L8s", capacity: 2, tier: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

kustoClustersCreateOrUpdate().catch(console.error);
