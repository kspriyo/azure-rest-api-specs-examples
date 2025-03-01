const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Polls an Endpoint operation.
 *
 * @summary Polls an Endpoint operation.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/getLogs.json
 */
async function getOnlineDeploymentLogs() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspace123";
  const endpointName = "testEndpoint";
  const deploymentName = "testDeployment";
  const body = {
    containerType: "StorageInitializer",
    tail: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.onlineDeployments.getLogs(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName,
    body
  );
  console.log(result);
}

getOnlineDeploymentLogs().catch(console.error);
