const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Recommended tasks that will help improve the security of the subscription proactively
 *
 * @summary Recommended tasks that will help improve the security of the subscription proactively
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/UpdateTaskSubscriptionLocation_example.json
 */
async function changeSecurityRecommendationTaskState() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const taskName = "62609ee7-d0a5-8616-9fe4-1df5cca7758d";
  const taskUpdateActionType = "Dismiss";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.tasks.updateSubscriptionLevelTaskState(
    ascLocation,
    taskName,
    taskUpdateActionType
  );
  console.log(result);
}
