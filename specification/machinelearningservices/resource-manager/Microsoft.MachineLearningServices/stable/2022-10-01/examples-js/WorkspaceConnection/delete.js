const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/WorkspaceConnection/delete.json
 */
async function deleteWorkspaceConnection() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup-1";
  const workspaceName = "workspace-1";
  const connectionName = "connection-1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaceConnections.delete(
    resourceGroupName,
    workspaceName,
    connectionName
  );
  console.log(result);
}

deleteWorkspaceConnection().catch(console.error);
