const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of all supported Security Solutions for subscription and location.
 *
 * @summary Gets list of all supported Security Solutions for subscription and location.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutionsReferenceData/GetSecuritySolutionsReferenceDataSubscriptionLocation_example.json
 */
async function getSecuritySolutionsFromASecurityDataLocation() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securitySolutionsReferenceDataOperations.listByHomeRegion(
    ascLocation
  );
  console.log(result);
}
