const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to A live event in Stopped or StandBy state will be in Running state after the start operation completes.
 *
 * @summary A live event in Stopped or StandBy state will be in Running state after the start operation completes.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-start.json
 */
async function startALiveEvent() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveEvents.beginStartAndWait(
    resourceGroupName,
    accountName,
    liveEventName
  );
  console.log(result);
}

startALiveEvent().catch(console.error);
