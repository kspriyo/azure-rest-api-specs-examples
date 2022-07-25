const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove workload classifier of a Sql pool's workload group.
 *
 * @summary Remove workload classifier of a Sql pool's workload group.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroupWorkloadClassifer.json
 */
async function deleteAWorkloadClassifierOfASqlAnalyticsPoolWorkloadGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const workloadGroupName = "wlm_workloadgroup";
  const workloadClassifierName = "wlm_workloadclassifier";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolWorkloadClassifier.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    workloadGroupName,
    workloadClassifierName
  );
  console.log(result);
}

deleteAWorkloadClassifierOfASqlAnalyticsPoolWorkloadGroup().catch(console.error);
