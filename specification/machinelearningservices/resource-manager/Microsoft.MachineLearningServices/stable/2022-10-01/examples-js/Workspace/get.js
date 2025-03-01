const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified machine learning workspace.
 *
 * @summary Gets the properties of the specified machine learning workspace.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/get.json
 */
async function getWorkspace() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-1234";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaces.get(resourceGroupName, workspaceName);
  console.log(result);
}

getWorkspace().catch(console.error);
