const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a cloud service.
 *
 * @summary Deletes a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudService_Delete.json
 */
async function deleteCloudService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.beginDeleteAndWait(resourceGroupName, cloudServiceName);
  console.log(result);
}

deleteCloudService().catch(console.error);
