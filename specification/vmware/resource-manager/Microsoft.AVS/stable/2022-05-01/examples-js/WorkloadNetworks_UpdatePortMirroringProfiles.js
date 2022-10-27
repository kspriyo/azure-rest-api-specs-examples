const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a port mirroring profile by id in a private cloud workload network.
 *
 * @summary Create or update a port mirroring profile by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_UpdatePortMirroringProfiles.json
 */
async function workloadNetworksUpdatePortMirroring() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const portMirroringId = "portMirroring1";
  const workloadNetworkPortMirroring = {
    destination: "vmGroup2",
    direction: "BIDIRECTIONAL",
    revision: 1,
    source: "vmGroup1",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginUpdatePortMirroringAndWait(
    resourceGroupName,
    privateCloudName,
    portMirroringId,
    workloadNetworkPortMirroring
  );
  console.log(result);
}

workloadNetworksUpdatePortMirroring().catch(console.error);
