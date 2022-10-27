const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a DNS zone by id in a private cloud workload network.
 *
 * @summary Create a DNS zone by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_CreateDnsZones.json
 */
async function workloadNetworksCreateDnsZone() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const dnsZoneId = "dnsZone1";
  const workloadNetworkDnsZone = {
    displayName: "dnsZone1",
    dnsServerIps: ["1.1.1.1"],
    domain: [],
    revision: 1,
    sourceIp: "8.8.8.8",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginCreateDnsZoneAndWait(
    resourceGroupName,
    privateCloudName,
    dnsZoneId,
    workloadNetworkDnsZone
  );
  console.log(result);
}

workloadNetworksCreateDnsZone().catch(console.error);
