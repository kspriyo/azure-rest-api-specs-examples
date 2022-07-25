const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a workspace managed sql server's security alert policy.
 *
 * @summary Get a workspace managed sql server's security alert policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerSecurityAlertPolicy.json
 */
async function getWorkspaceManagedSqlServerSecurityAlertPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const securityAlertPolicyName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedSqlServerSecurityAlertPolicy.get(
    resourceGroupName,
    workspaceName,
    securityAlertPolicyName
  );
  console.log(result);
}

getWorkspaceManagedSqlServerSecurityAlertPolicy().catch(console.error);
