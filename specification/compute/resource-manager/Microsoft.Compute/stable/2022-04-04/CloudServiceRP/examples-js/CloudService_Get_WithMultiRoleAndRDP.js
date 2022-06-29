const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Display information about a cloud service.
 *
 * @summary Display information about a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudService_Get_WithMultiRoleAndRDP.json
 */
async function getCloudServiceWithMultipleRolesAndRdpExtension() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.get(resourceGroupName, cloudServiceName);
  console.log(result);
}

getCloudServiceWithMultipleRolesAndRdpExtension().catch(console.error);
