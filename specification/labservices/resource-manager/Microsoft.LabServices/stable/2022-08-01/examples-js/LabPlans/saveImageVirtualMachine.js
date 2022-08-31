const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Saves an image from a lab VM to the attached shared image gallery.
 *
 * @summary Saves an image from a lab VM to the attached shared image gallery.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/saveImageVirtualMachine.json
 */
async function saveImageVirtualMachine() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labPlanName = "testlabplan";
  const body = {
    name: "Test Image",
    labVirtualMachineId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labs/testlab/virtualMachines/template",
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labPlans.beginSaveImageAndWait(resourceGroupName, labPlanName, body);
  console.log(result);
}

saveImageVirtualMachine().catch(console.error);
