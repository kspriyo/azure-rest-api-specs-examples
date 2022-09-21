const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the cluster name is valid and is not already in use.
 *
 * @summary Checks that the cluster name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClustersCheckNameAvailability.json
 */
async function kustoClustersCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const location = "westus";
  const clusterName = {
    name: "kustoCluster",
    type: "Microsoft.Kusto/clusters",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.checkNameAvailability(location, clusterName);
  console.log(result);
}

kustoClustersCheckNameAvailability().catch(console.error);
