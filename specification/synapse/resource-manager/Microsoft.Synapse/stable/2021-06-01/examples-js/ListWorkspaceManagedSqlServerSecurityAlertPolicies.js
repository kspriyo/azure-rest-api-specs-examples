const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get workspace managed sql server's threat detection policies.
 *
 * @summary Get workspace managed sql server's threat detection policies.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerSecurityAlertPolicies.json
 */
async function getWorkspaceManagedSqlServerSecurityAlertPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceManagedSqlServerSecurityAlertPolicy.list(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getWorkspaceManagedSqlServerSecurityAlertPolicy().catch(console.error);
