const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the role instances in the specified update domain.
 *
 * @summary Updates the role instances in the specified update domain.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudServiceUpdateDomain_Update.json
 */
async function updateCloudServiceToSpecifiedDomain() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const updateDomain = 1;
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServicesUpdateDomain.beginWalkUpdateDomainAndWait(
    resourceGroupName,
    cloudServiceName,
    updateDomain
  );
  console.log(result);
}

updateCloudServiceToSpecifiedDomain().catch(console.error);
