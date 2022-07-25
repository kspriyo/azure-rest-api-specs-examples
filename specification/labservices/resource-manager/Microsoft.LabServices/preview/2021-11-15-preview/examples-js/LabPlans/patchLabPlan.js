const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update a Lab Plan resource.
 *
 * @summary Operation to update a Lab Plan resource.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/patchLabPlan.json
 */
async function patchLabPlan() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labPlanName = "testlabplan";
  const body = {
    defaultConnectionProfile: {
      clientRdpAccess: "Public",
      clientSshAccess: "Public",
      webRdpAccess: "None",
      webSshAccess: "None",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labPlans.beginUpdateAndWait(resourceGroupName, labPlanName, body);
  console.log(result);
}

patchLabPlan().catch(console.error);
