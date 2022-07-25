const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the environment with the specified name in the specified subscription and resource group.
 *
 * @summary Updates the environment with the specified name in the specified subscription and resource group.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EnvironmentsPatchTags.json
 */
async function environmentsUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const environmentUpdateParameters = {
    tags: { someTag: "someTagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.environments.beginUpdateAndWait(
    resourceGroupName,
    environmentName,
    environmentUpdateParameters
  );
  console.log(result);
}

environmentsUpdate().catch(console.error);
