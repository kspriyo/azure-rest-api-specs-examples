const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a firewall rule
 *
 * @summary Get a firewall rule
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetIpFirewallRule.json
 */
async function createAnIPFirewallRule() {
  const subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const ruleName = "ExampleIpFirewallRule";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.ipFirewallRules.get(resourceGroupName, workspaceName, ruleName);
  console.log(result);
}

createAnIPFirewallRule().catch(console.error);
