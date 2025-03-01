const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a valid AAD token for an Endpoint using AMLToken-based authentication.
 *
 * @summary Retrieve a valid AAD token for an Endpoint using AMLToken-based authentication.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/getToken.json
 */
async function getTokenOnlineEndpoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.onlineEndpoints.getToken(
    resourceGroupName,
    workspaceName,
    endpointName
  );
  console.log(result);
}

getTokenOnlineEndpoint().catch(console.error);
