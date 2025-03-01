const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
 *
 * @summary Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/listKeys.json
 */
async function listAksComputeKeys() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.computeOperations.listKeys(
    resourceGroupName,
    workspaceName,
    computeName
  );
  console.log(result);
}

listAksComputeKeys().catch(console.error);
