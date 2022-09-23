const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the quota (service limits) of this resource to the requested value.
  • To get the quota information for specific resource, send a GET request.
  • To increase the quota, update the limit field from the GET response to a new value.
  • To update the quota value, submit the JSON response to the quota request API to update the quota.
  • To update the quota. use the PATCH operation.
 *
 * @summary Update the quota (service limits) of this resource to the requested value.
  • To get the quota information for specific resource, send a GET request.
  • To increase the quota, update the limit field from the GET response to a new value.
  • To update the quota value, submit the JSON response to the quota request API to update the quota.
  • To update the quota. use the PATCH operation.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/patchComputeQuotaRequest.json
 */
async function quotasRequestPatchForCompute() {
  const subscriptionId = "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3";
  const providerId = "Microsoft.Compute";
  const location = "eastus";
  const resourceName = "standardFSv2Family";
  const createQuotaRequest = {
    properties: {
      name: { value: "standardFSv2Family" },
      limit: 200,
      unit: "Count",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.quota.beginUpdateAndWait(
    subscriptionId,
    providerId,
    location,
    resourceName,
    createQuotaRequest
  );
  console.log(result);
}

quotasRequestPatchForCompute().catch(console.error);
