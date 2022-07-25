const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Sql pool's connection policy, which is used with table auditing.
 *
 * @summary Get a Sql pool's connection policy, which is used with table auditing.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolConnectionPolicy.json
 */
async function getAConnectionPolicyOfASqlAnalyticsPool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-6852";
  const workspaceName = "blobauditingtest-2080";
  const sqlPoolName = "testdb";
  const connectionPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolConnectionPolicies.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    connectionPolicyName
  );
  console.log(result);
}

getAConnectionPolicyOfASqlAnalyticsPool().catch(console.error);
