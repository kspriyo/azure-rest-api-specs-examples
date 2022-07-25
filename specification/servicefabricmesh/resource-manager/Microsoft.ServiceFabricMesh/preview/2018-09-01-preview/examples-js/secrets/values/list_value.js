const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the decrypted value of the specified named value of the secret resource. This is a privileged operation.
 *
 * @summary Lists the decrypted value of the specified named value of the secret resource. This is a privileged operation.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/values/list_value.json
 */
async function listSecretValue() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const secretValueResourceName = "v1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.secretValueOperations.listValue(
    resourceGroupName,
    secretResourceName,
    secretValueResourceName
  );
  console.log(result);
}

listSecretValue().catch(console.error);
