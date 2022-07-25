const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the reference data set with the specified name in the specified subscription, resource group, and environment.
 *
 * @summary Updates the reference data set with the specified name in the specified subscription, resource group, and environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/ReferenceDataSetsPatchTags.json
 */
async function referenceDataSetsUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const referenceDataSetName = "rds1";
  const referenceDataSetUpdateParameters = {
    tags: { someKey: "someValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.referenceDataSets.update(
    resourceGroupName,
    environmentName,
    referenceDataSetName,
    referenceDataSetUpdateParameters
  );
  console.log(result);
}

referenceDataSetsUpdate().catch(console.error);
