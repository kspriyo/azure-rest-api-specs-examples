const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if a DigitalTwinsInstance name is available.
 *
 * @summary Check if a DigitalTwinsInstance name is available.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsCheckNameAvailability_example.json
 */
async function checkNameAvailability() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const location = "WestUS2";
  const digitalTwinsInstanceCheckName = {
    name: "myadtinstance",
    type: "Microsoft.DigitalTwins/digitalTwinsInstances",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.digitalTwins.checkNameAvailability(
    location,
    digitalTwinsInstanceCheckName
  );
  console.log(result);
}

checkNameAvailability().catch(console.error);
