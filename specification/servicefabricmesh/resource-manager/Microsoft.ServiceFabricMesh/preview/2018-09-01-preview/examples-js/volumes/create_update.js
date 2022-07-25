const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a volume resource with the specified name, description and properties. If a volume resource with the same name exists, then it is updated with the specified description and properties.
 *
 * @summary Creates a volume resource with the specified name, description and properties. If a volume resource with the same name exists, then it is updated with the specified description and properties.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/volumes/create_update.json
 */
async function createOrUpdateVolume() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const volumeResourceName = "sampleVolume";
  const volumeResourceDescription = {
    description: "Service Fabric Mesh sample volume.",
    azureFileParameters: {
      accountKey: "provide-account-key-here",
      accountName: "sbzdemoaccount",
      shareName: "sharel",
    },
    location: "EastUS",
    provider: "SFAzureFile",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.volume.create(
    resourceGroupName,
    volumeResourceName,
    volumeResourceDescription
  );
  console.log(result);
}

createOrUpdateVolume().catch(console.error);
