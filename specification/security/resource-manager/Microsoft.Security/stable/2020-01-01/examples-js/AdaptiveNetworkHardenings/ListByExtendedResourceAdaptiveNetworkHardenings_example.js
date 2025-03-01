const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Adaptive Network Hardenings resources in scope of an extended resource.
 *
 * @summary Gets a list of Adaptive Network Hardenings resources in scope of an extended resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/ListByExtendedResourceAdaptiveNetworkHardenings_example.json
 */
async function listAdaptiveNetworkHardeningsResourcesOfAnExtendedResource() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "rg1";
  const resourceNamespace = "Microsoft.Compute";
  const resourceType = "virtualMachines";
  const resourceName = "vm1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.adaptiveNetworkHardenings.listByExtendedResource(
    resourceGroupName,
    resourceNamespace,
    resourceType,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
