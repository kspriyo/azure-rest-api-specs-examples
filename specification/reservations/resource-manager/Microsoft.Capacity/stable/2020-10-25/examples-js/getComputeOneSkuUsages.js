const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the current quota (service limit) and usage of a resource. You can use the response from the GET operation to submit quota update request.
 *
 * @summary Get the current quota (service limit) and usage of a resource. You can use the response from the GET operation to submit quota update request.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeOneSkuUsages.json
 */
async function quotasRequestForCompute() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const providerId = "Microsoft.Compute";
  const location = "eastus";
  const resourceName = "standardNDSFamily";
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.quota.get(subscriptionId, providerId, location, resourceName);
  console.log(result);
}

quotasRequestForCompute().catch(console.error);
