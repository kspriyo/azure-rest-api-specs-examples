const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to check for resource name availability.
A name is available if no other resource exists that has the same SubscriptionId, Resource Name and Type
or if one or more such resources exist, each of these must be GC'd and their time of deletion be more than 24 Hours Ago
 *
 * @summary API to check for resource name availability.
A name is available if no other resource exists that has the same SubscriptionId, Resource Name and Type
or if one or more such resources exist, each of these must be GC'd and their time of deletion be more than 24 Hours Ago
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/CheckNameAvailability_NotAvailable.json
 */
async function availabilityStatusOfResourceNameWhenResourceWithSameNameTypeAndSubscriptionExists() {
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = "resGroupBar";
  const location = "westus";
  const input = {
    name: "swaggerExample2",
    type: "Microsoft.RecoveryServices/Vaults",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.recoveryServices.checkNameAvailability(
    resourceGroupName,
    location,
    input
  );
  console.log(result);
}

availabilityStatusOfResourceNameWhenResourceWithSameNameTypeAndSubscriptionExists().catch(
  console.error
);
