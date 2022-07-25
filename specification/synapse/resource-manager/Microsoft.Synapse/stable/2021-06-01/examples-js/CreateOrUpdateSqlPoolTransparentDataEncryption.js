const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Sql pool's transparent data encryption configuration.
 *
 * @summary Creates or updates a Sql pool's transparent data encryption configuration.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolTransparentDataEncryption.json
 */
async function createOrUpdateASqlPoolTransparentDataEncryptionConfiguration() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const transparentDataEncryptionName = "current";
  const parameters = { status: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolTransparentDataEncryptions.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    transparentDataEncryptionName,
    parameters
  );
  console.log(result);
}

createOrUpdateASqlPoolTransparentDataEncryptionConfiguration().catch(console.error);
