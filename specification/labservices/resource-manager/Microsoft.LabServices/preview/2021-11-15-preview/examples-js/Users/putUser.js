const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to create or update a lab user.
 *
 * @summary Operation to create or update a lab user.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/putUser.json
 */
async function putUser() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const userName = "testuser";
  const body = {
    additionalUsageQuota: "20:00",
    email: "testuser@contoso.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.users.beginCreateOrUpdateAndWait(
    resourceGroupName,
    labName,
    userName,
    body
  );
  console.log(result);
}

putUser().catch(console.error);
