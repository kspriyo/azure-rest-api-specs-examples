const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to For the specified Azure region (location), get the details and status of the quota request by the quota request ID for the resources of the resource provider. The PUT request for the quota (service limit) returns a response with the requestId parameter.
 *
 * @summary For the specified Azure region (location), get the details and status of the quota request by the quota request ID for the resources of the resource provider. The PUT request for the quota (service limit) returns a response with the requestId parameter.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getQuotaRequestStatusFailed.json
 */
async function quotaRequestFailed() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const providerId = "Microsoft.Compute";
  const location = "eastus";
  const id = "2B5C8515-37D8-4B6A-879B-CD641A2CF605";
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.quotaRequestStatus.get(subscriptionId, providerId, location, id);
  console.log(result);
}

quotaRequestFailed().catch(console.error);
