const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a cloud link in a private cloud
 *
 * @summary Delete a cloud link in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/CloudLinks_Delete.json
 */
async function cloudLinksDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const cloudLinkName = "cloudLink1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.cloudLinks.beginDeleteAndWait(
    resourceGroupName,
    privateCloudName,
    cloudLinkName
  );
  console.log(result);
}

cloudLinksDelete().catch(console.error);
