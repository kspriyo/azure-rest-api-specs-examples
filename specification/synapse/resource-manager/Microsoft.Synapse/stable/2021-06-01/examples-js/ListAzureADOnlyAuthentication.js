const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Azure Active Directory only authentication property for a workspace
 *
 * @summary Gets a list of Azure Active Directory only authentication property for a workspace
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListAzureADOnlyAuthentication.json
 */
async function getAListOfAzureActiveDirectoryOnlyAuthenticationProperty() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-6852";
  const workspaceName = "workspace-2080";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureADOnlyAuthentications.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfAzureActiveDirectoryOnlyAuthenticationProperty().catch(console.error);
