const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Kusto cluster database principalAssignment.
 *
 * @summary Creates a Kusto cluster database principalAssignment.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDatabasePrincipalAssignmentsCreateOrUpdate.json
 */
async function kustoDatabasePrincipalAssignmentsCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "Kustodatabase8";
  const principalAssignmentName = "kustoprincipal1";
  const parameters = {
    principalId: "87654321-1234-1234-1234-123456789123",
    principalType: "App",
    role: "Admin",
    tenantId: "12345678-1234-1234-1234-123456789123",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databasePrincipalAssignments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    principalAssignmentName,
    parameters
  );
  console.log(result);
}

kustoDatabasePrincipalAssignmentsCreateOrUpdate().catch(console.error);
