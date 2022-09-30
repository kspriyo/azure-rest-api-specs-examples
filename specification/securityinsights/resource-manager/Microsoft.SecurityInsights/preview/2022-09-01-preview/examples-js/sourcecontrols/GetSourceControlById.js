const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a source control byt its identifier.
 *
 * @summary Gets a source control byt its identifier.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/GetSourceControlById.json
 */
async function getASourceControl() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const sourceControlId = "789e0c1f-4a3d-43ad-809c-e713b677b04a";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.sourceControls.get(resourceGroupName, workspaceName, sourceControlId);
  console.log(result);
}

getASourceControl().catch(console.error);
