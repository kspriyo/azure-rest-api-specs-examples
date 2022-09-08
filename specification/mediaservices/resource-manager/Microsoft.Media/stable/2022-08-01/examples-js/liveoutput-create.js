const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new live output.
 *
 * @summary Creates a new live output.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveoutput-create.json
 */
async function createALiveOutput() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const liveOutputName = "myLiveOutput1";
  const parameters = {
    description: "test live output 1",
    archiveWindowLength: "PT5M",
    assetName: "6f3264f5-a189-48b4-a29a-a40f22575212",
    hls: { fragmentsPerTsSegment: 5 },
    manifestName: "testmanifest",
    rewindWindowLength: "PT4M",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveOutputs.beginCreateAndWait(
    resourceGroupName,
    accountName,
    liveEventName,
    liveOutputName,
    parameters
  );
  console.log(result);
}

createALiveOutput().catch(console.error);
