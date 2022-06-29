const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of all role instances in a cloud service. Use nextLink property in the response to get the next page of role instances. Do this till nextLink is null to fetch all the role instances.
 *
 * @summary Gets the list of all role instances in a cloud service. Use nextLink property in the response to get the next page of role instances. Do this till nextLink is null to fetch all the role instances.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudServiceRolesInstance_List.json
 */
async function listRoleInstancesInACloudService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cloudServiceRoleInstances.list(
    resourceGroupName,
    cloudServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRoleInstancesInACloudService().catch(console.error);
