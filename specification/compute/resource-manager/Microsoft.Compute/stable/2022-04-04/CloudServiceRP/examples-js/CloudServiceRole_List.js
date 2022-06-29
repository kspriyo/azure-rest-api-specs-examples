const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all roles in a cloud service. Use nextLink property in the response to get the next page of roles. Do this till nextLink is null to fetch all the roles.
 *
 * @summary Gets a list of all roles in a cloud service. Use nextLink property in the response to get the next page of roles. Do this till nextLink is null to fetch all the roles.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudServiceRole_List.json
 */
async function listRolesInACloudService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cloudServiceRoles.list(resourceGroupName, cloudServiceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRolesInACloudService().catch(console.error);
