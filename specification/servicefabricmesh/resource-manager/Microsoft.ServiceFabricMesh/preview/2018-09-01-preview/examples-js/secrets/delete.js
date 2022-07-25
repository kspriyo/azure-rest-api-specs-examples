const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the secret resource identified by the name.
 *
 * @summary Deletes the secret resource identified by the name.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/delete.json
 */
async function deleteSecret() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.secret.delete(resourceGroupName, secretResourceName);
  console.log(result);
}

deleteSecret().catch(console.error);
