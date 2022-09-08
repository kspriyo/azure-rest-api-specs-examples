const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Content Key Policy in the Media Services account
 *
 * @summary Deletes a Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-delete.json
 */
async function deleteAKeyPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithPlayReadyOptionAndOpenRestriction";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.contentKeyPolicies.delete(
    resourceGroupName,
    accountName,
    contentKeyPolicyName
  );
  console.log(result);
}

deleteAKeyPolicy().catch(console.error);
