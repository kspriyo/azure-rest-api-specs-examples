const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a deleted sql pool that can be restored
 *
 * @summary Gets a deleted sql pool that can be restored
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlPoolGet.json
 */
async function getARestorableDroppedSqlPool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "restorabledroppeddatabasetest-1257";
  const workspaceName = "restorabledroppeddatabasetest-2389";
  const restorableDroppedSqlPoolId = "restorabledroppeddatabasetest-7654,131403269876900000";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.restorableDroppedSqlPools.get(
    resourceGroupName,
    workspaceName,
    restorableDroppedSqlPoolId
  );
  console.log(result);
}

getARestorableDroppedSqlPool().catch(console.error);
