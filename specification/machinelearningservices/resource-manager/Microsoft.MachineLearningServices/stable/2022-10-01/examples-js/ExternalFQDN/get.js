const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Called by Client (Portal, CLI, etc) to get a list of all external outbound dependencies (FQDNs) programmatically.
 *
 * @summary Called by Client (Portal, CLI, etc) to get a list of all external outbound dependencies (FQDNs) programmatically.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ExternalFQDN/get.json
 */
async function listOutboundNetworkDependenciesEndpoints() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-1234";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaces.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

listOutboundNetworkDependenciesEndpoints().catch(console.error);
