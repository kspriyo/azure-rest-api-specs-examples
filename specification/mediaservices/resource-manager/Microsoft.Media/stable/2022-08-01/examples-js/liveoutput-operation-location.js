const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Live Output operation status.
 *
 * @summary Get a Live Output operation status.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveoutput-operation-location.json
 */
async function getTheLiveOutputOperationStatus() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const liveOutputName = "myLiveOutput1";
  const operationId = "62e4d893-d233-4005-988e-a428d9f77076";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveOutputs.operationLocation(
    resourceGroupName,
    accountName,
    liveEventName,
    liveOutputName,
    operationId
  );
  console.log(result);
}

getTheLiveOutputOperationStatus().catch(console.error);
