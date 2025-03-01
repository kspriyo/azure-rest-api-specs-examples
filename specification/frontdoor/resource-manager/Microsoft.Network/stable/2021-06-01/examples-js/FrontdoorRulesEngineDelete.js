const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Rules Engine Configuration with the specified parameters.
 *
 * @summary Deletes an existing Rules Engine Configuration with the specified parameters.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineDelete.json
 */
async function deleteRulesEngineConfiguration() {
  const subscriptionId = process.env["FRONTDOOR_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["FRONTDOOR_RESOURCE_GROUP"] || "rg1";
  const frontDoorName = "frontDoor1";
  const rulesEngineName = "rulesEngine1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.rulesEngines.beginDeleteAndWait(
    resourceGroupName,
    frontDoorName,
    rulesEngineName
  );
  console.log(result);
}
