const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Asset Filter associated with the specified Asset.
 *
 * @summary Updates an existing Asset Filter associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assetFilters-update.json
 */
async function updateAnAssetFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const filterName = "assetFilterWithTimeWindowAndTrack";
  const parameters = {
    firstQuality: { bitrate: 128000 },
    presentationTimeRange: {
      endTimestamp: 170000000,
      forceEndTimestamp: false,
      liveBackoffDuration: 0,
      presentationWindowDuration: 9223372036854775000,
      startTimestamp: 10,
      timescale: 10000000,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assetFilters.update(
    resourceGroupName,
    accountName,
    assetName,
    filterName,
    parameters
  );
  console.log(result);
}

updateAnAssetFilter().catch(console.error);
