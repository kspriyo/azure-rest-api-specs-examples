const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Kusto cluster principalAssignment.
 *
 * @summary Deletes a Kusto cluster principalAssignment.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClusterPrincipalAssignmentsDelete.json
 */
async function kustoClusterPrincipalAssignmentsDelete() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const principalAssignmentName = "kustoprincipal1";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusterPrincipalAssignments.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    principalAssignmentName
  );
  console.log(result);
}

kustoClusterPrincipalAssignmentsDelete().catch(console.error);
