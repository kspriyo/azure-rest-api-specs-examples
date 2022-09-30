const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get threat intelligence indicators metrics (Indicators counts by Type, Threat Type, Source).
 *
 * @summary Get threat intelligence indicators metrics (Indicators counts by Type, Threat Type, Source).
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/CollectThreatIntelligenceMetrics.json
 */
async function getThreatIntelligenceIndicatorsMetrics() {
  const subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.threatIntelligenceIndicatorMetrics.list(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

getThreatIntelligenceIndicatorsMetrics().catch(console.error);
