const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the alert's state
 *
 * @summary Update the alert's state
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/UpdateAlertResourceGroupLocation_resolve_example.json
 */
async function updateSecurityAlertStateOnAResourceGroupFromASecurityDataLocation() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg2";
  const ascLocation = "westeurope";
  const alertName = "2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.updateResourceGroupLevelStateToResolve(
    resourceGroupName,
    ascLocation,
    alertName
  );
  console.log(result);
}
