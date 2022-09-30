const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a source control.
 *
 * @summary Delete a source control.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/DeleteSourceControl.json
 */
async function deleteASourceControl() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const sourceControlId = "789e0c1f-4a3d-43ad-809c-e713b677b04a";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.sourceControls.delete(
    resourceGroupName,
    workspaceName,
    sourceControlId
  );
  console.log(result);
}

deleteASourceControl().catch(console.error);
