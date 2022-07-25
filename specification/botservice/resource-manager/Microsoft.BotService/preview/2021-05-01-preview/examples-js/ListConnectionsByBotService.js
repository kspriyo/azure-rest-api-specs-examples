const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all the Connection Settings registered to a particular BotService resource
 *
 * @summary Returns all the Connection Settings registered to a particular BotService resource
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListConnectionsByBotService.json
 */
async function listConnectionSettings() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.botConnection.listByBotService(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConnectionSettings().catch(console.error);
