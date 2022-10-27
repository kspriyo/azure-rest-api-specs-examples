const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a vm group by id in a private cloud workload network.
 *
 * @summary Delete a vm group by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_DeleteVMGroups.json
 */
async function workloadNetworksDeleteVMGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const vmGroupId = "vmGroup1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginDeleteVMGroupAndWait(
    resourceGroupName,
    vmGroupId,
    privateCloudName
  );
  console.log(result);
}

workloadNetworksDeleteVMGroup().catch(console.error);
