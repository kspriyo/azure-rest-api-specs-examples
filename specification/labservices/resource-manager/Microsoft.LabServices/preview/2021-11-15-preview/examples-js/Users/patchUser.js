const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update a lab user.
 *
 * @summary Operation to update a lab user.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/patchUser.json
 */
async function patchUser() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const userName = "testuser";
  const body = { additionalUsageQuota: "20:00" };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.users.beginUpdateAndWait(resourceGroupName, labName, userName, body);
  console.log(result);
}

patchUser().catch(console.error);
