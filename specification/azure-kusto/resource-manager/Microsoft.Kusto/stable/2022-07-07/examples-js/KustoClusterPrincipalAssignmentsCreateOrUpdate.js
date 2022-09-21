const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Kusto cluster principalAssignment.
 *
 * @summary Create a Kusto cluster principalAssignment.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClusterPrincipalAssignmentsCreateOrUpdate.json
 */
async function kustoClusterPrincipalAssignmentsCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const principalAssignmentName = "kustoprincipal1";
  const parameters = {
    principalId: "87654321-1234-1234-1234-123456789123",
    principalType: "App",
    role: "AllDatabasesAdmin",
    tenantId: "12345678-1234-1234-1234-123456789123",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusterPrincipalAssignments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    principalAssignmentName,
    parameters
  );
  console.log(result);
}

kustoClusterPrincipalAssignmentsCreateOrUpdate().catch(console.error);
