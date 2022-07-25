const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update a Sql pool's workload group.
 *
 * @summary Create Or Update a Sql pool's workload group.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadGroupMax.json
 */
async function createAWorkloadGroupWithAllPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const workloadGroupName = "smallrc";
  const parameters = {
    importance: "normal",
    maxResourcePercent: 100,
    maxResourcePercentPerRequest: 3,
    minResourcePercent: 0,
    minResourcePercentPerRequest: 3,
    queryExecutionTimeout: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolWorkloadGroup.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    workloadGroupName,
    parameters
  );
  console.log(result);
}

createAWorkloadGroupWithAllPropertiesSpecified().catch(console.error);
