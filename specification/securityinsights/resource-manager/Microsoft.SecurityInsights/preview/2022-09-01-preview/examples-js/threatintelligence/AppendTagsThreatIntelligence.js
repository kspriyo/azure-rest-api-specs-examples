const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Append tags to a threat intelligence indicator.
 *
 * @summary Append tags to a threat intelligence indicator.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/AppendTagsThreatIntelligence.json
 */
async function appendTagsToAThreatIntelligenceIndicator() {
  const subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const name = "d9cd6f0b-96b9-3984-17cd-a779d1e15a93";
  const threatIntelligenceAppendTags = {
    threatIntelligenceTags: ["tag1", "tag2"],
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.threatIntelligenceIndicator.appendTags(
    resourceGroupName,
    workspaceName,
    name,
    threatIntelligenceAppendTags
  );
  console.log(result);
}

appendTagsToAThreatIntelligenceIndicator().catch(console.error);
