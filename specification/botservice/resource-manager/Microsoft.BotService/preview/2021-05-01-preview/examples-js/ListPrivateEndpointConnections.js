const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the private endpoint connections associated with the Bot.
 *
 * @summary List all the private endpoint connections associated with the Bot.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListPrivateEndpointConnections.json
 */
async function listPrivateEndpointConnections() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const resourceName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPrivateEndpointConnections().catch(console.error);
