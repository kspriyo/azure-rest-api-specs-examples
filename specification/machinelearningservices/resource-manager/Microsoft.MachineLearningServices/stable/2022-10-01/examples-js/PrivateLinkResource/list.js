const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a workspace.
 *
 * @summary Gets the private link resources that need to be created for a workspace.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/PrivateLinkResource/list.json
 */
async function workspaceListPrivateLinkResources() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "rg-1234";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.privateLinkResources.list(resourceGroupName, workspaceName);
  console.log(result);
}

workspaceListPrivateLinkResources().catch(console.error);
