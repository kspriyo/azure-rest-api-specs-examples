const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of endpoints that VNET Injected Workspace calls Azure Databricks Control Plane. You must configure outbound access with these endpoints. For more information, see https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/udr
 *
 * @summary Gets the list of endpoints that VNET Injected Workspace calls Azure Databricks Control Plane. You must configure outbound access with these endpoints. For more information, see https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/udr
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/OutboundNetworkDependenciesEndpointsList.json
 */
async function listOutboundNetworkDependenciesEndpointsByWorkspace() {
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.outboundNetworkDependenciesEndpoints.list(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

listOutboundNetworkDependenciesEndpointsByWorkspace().catch(console.error);
