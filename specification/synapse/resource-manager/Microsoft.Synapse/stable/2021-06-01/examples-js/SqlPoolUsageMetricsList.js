const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets SQL pool usages.
 *
 * @summary Gets SQL pool usages.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolUsageMetricsList.json
 */
async function listTheUsagesOfASqlAnalyticsPool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6730";
  const workspaceName = "sqlcrudtest-9007";
  const sqlPoolName = "3481";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlPoolUsages.list(resourceGroupName, workspaceName, sqlPoolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheUsagesOfASqlAnalyticsPool().catch(console.error);
