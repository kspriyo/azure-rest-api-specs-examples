const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Azure Active Directory only authentication property
 *
 * @summary Gets a Azure Active Directory only authentication property
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAzureADOnlyAuthentication.json
 */
async function getAzureActiveDirectoryOnlyAuthenticationProperty() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-6852";
  const workspaceName = "workspace-2080";
  const azureADOnlyAuthenticationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.azureADOnlyAuthentications.get(
    resourceGroupName,
    workspaceName,
    azureADOnlyAuthenticationName
  );
  console.log(result);
}

getAzureActiveDirectoryOnlyAuthenticationProperty().catch(console.error);
