const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Kusto cluster database principalAssignment.
 *
 * @summary Gets a Kusto cluster database principalAssignment.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDatabasePrincipalAssignmentsGet.json
 */
async function kustoDatabasePrincipalAssignmentsGet() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "Kustodatabase8";
  const principalAssignmentName = "kustoprincipal1";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databasePrincipalAssignments.get(
    resourceGroupName,
    clusterName,
    databaseName,
    principalAssignmentName
  );
  console.log(result);
}

kustoDatabasePrincipalAssignmentsGet().catch(console.error);
