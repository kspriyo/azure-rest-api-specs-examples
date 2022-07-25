const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements AvailabilitySet GET method.
 *
 * @summary Implements AvailabilitySet GET method.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetAvailabilitySet.json
 */
async function getAvailabilitySet() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const availabilitySetName = "HRAvailabilitySet";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.availabilitySets.get(resourceGroupName, availabilitySetName);
  console.log(result);
}

getAvailabilitySet().catch(console.error);
