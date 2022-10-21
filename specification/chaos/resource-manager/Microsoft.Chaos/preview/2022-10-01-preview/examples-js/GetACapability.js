const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Capability resource that extends a Target resource.
 *
 * @summary Get a Capability resource that extends a Target resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/GetACapability.json
 */
async function getACapabilityThatExtendsAVirtualMachineTargetResource() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const parentProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResourceName = "exampleVM";
  const targetName = "Microsoft-VirtualMachine";
  const capabilityName = "Shutdown-1.0";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.capabilities.get(
    resourceGroupName,
    parentProviderNamespace,
    parentResourceType,
    parentResourceName,
    targetName,
    capabilityName
  );
  console.log(result);
}

getACapabilityThatExtendsAVirtualMachineTargetResource().catch(console.error);
