const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a segment by id in a private cloud workload network.
 *
 * @summary Create a segment by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_CreateSegments.json
 */
async function workloadNetworksCreateSegments() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const segmentId = "segment1";
  const workloadNetworkSegment = {
    connectedGateway: "/infra/tier-1s/gateway",
    displayName: "segment1",
    revision: 1,
    subnet: {
      dhcpRanges: ["40.20.0.0-40.20.0.1"],
      gatewayAddress: "40.20.20.20/16",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginCreateSegmentsAndWait(
    resourceGroupName,
    privateCloudName,
    segmentId,
    workloadNetworkSegment
  );
  console.log(result);
}

workloadNetworksCreateSegments().catch(console.error);
