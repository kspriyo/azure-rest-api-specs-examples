const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List workspace managed sql server's extended blob auditing policies.
 *
 * @summary List workspace managed sql server's extended blob auditing policies.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerExtendedBlobAuditingSettings.json
 */
async function getWorkspaceManagedSqlServerExtendedBlobAuditingSettings() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceManagedSqlServerExtendedBlobAuditingPolicies.listByWorkspace(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getWorkspaceManagedSqlServerExtendedBlobAuditingSettings().catch(console.error);
