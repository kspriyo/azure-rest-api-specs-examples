const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete Batch Inference Endpoint (asynchronous).
 *
 * @summary Delete Batch Inference Endpoint (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/delete.json
 */
async function deleteBatchEndpoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup-1234";
  const workspaceName = "testworkspace";
  const endpointName = "testBatchEndpoint";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.batchEndpoints.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    endpointName
  );
  console.log(result);
}

deleteBatchEndpoint().catch(console.error);
