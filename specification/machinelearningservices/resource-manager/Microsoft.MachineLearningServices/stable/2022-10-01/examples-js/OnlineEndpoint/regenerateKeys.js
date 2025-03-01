const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
 *
 * @summary Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/regenerateKeys.json
 */
async function regenerateKeysOnlineEndpoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const body = {
    keyType: "Primary",
    keyValue: "string",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.onlineEndpoints.beginRegenerateKeysAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    body
  );
  console.log(result);
}

regenerateKeysOnlineEndpoint().catch(console.error);
