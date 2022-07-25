const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a workspace managed sql server's blob auditing policy.
 *
 * @summary Get a workspace managed sql server's blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerBlobAuditingSettings.json
 */
async function getBlobAuditingSettingOfWorkspaceManagedSqlServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const blobAuditingPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedSqlServerBlobAuditingPolicies.get(
    resourceGroupName,
    workspaceName,
    blobAuditingPolicyName
  );
  console.log(result);
}

getBlobAuditingSettingOfWorkspaceManagedSqlServer().catch(console.error);
