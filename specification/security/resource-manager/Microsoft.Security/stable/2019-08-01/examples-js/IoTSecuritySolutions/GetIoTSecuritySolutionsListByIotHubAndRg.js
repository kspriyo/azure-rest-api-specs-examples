const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get the list IoT Security solutions organized by resource group.
 *
 * @summary Use this method to get the list IoT Security solutions organized by resource group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHubAndRg.json
 */
async function listIoTSecuritySolutionsByResourceGroupAndIoTHub() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "MyRg";
  const filter =
    'properties.iotHubs/any(i eq "/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")';
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotSecuritySolution.listByResourceGroup(
    resourceGroupName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
