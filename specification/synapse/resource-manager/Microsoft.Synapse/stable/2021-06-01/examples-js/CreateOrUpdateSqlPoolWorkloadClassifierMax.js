const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update workload classifier for a Sql pool's workload group.
 *
 * @summary Create Or Update workload classifier for a Sql pool's workload group.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadClassifierMax.json
 */
async function createAWorkloadClassifierWithAllPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const workloadGroupName = "wlm_workloadgroup";
  const workloadClassifierName = "wlm_workloadclassifier";
  const parameters = {
    context: "test_context",
    endTime: "14:00",
    importance: "high",
    label: "test_label",
    memberName: "dbo",
    startTime: "12:00",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolWorkloadClassifier.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    workloadGroupName,
    workloadClassifierName,
    parameters
  );
  console.log(result);
}

createAWorkloadClassifierWithAllPropertiesSpecified().catch(console.error);
