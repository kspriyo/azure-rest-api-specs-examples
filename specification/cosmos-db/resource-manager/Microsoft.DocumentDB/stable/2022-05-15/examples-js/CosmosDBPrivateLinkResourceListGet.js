const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a Cosmos DB account.
 *
 * @summary Gets the private link resources that need to be created for a Cosmos DB account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBPrivateLinkResourceListGet.json
 */
async function getsPrivateEndpointConnection() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByDatabaseAccount(
    resourceGroupName,
    accountName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsPrivateEndpointConnection().catch(console.error);
