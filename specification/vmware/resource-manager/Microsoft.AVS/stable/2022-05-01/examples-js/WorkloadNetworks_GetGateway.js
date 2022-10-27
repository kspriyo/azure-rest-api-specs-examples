const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a gateway by id in a private cloud workload network.
 *
 * @summary Get a gateway by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetGateway.json
 */
async function workloadNetworksGetGateway() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const gatewayId = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.getGateway(
    resourceGroupName,
    privateCloudName,
    gatewayId
  );
  console.log(result);
}

workloadNetworksGetGateway().catch(console.error);
