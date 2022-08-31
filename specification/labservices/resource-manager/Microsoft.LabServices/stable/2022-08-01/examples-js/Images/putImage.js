const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an image resource via PUT. Creating new resources via PUT will not function.
 *
 * @summary Updates an image resource via PUT. Creating new resources via PUT will not function.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Images/putImage.json
 */
async function putImage() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labPlanName = "testlabplan";
  const imageName = "image1";
  const body = { enabledState: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.images.createOrUpdate(
    resourceGroupName,
    labPlanName,
    imageName,
    body
  );
  console.log(result);
}

putImage().catch(console.error);
