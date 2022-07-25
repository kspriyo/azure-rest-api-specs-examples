const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the information about all replicas of a given service of an application. The information includes the runtime properties of the replica instance.
 *
 * @summary Gets the information about all replicas of a given service of an application. The information includes the runtime properties of the replica instance.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/replicas/list.json
 */
async function replicasGetAll() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const applicationResourceName = "sampleApplication";
  const serviceResourceName = "helloWorldService";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceReplica.list(
    resourceGroupName,
    applicationResourceName,
    serviceResourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

replicasGetAll().catch(console.error);
