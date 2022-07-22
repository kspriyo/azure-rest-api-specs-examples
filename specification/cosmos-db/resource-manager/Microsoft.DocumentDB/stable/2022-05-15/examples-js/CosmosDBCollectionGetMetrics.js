const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the metrics determined by the given filter for the given database account and collection.
 *
 * @summary Retrieves the metrics determined by the given filter for the given database account and collection.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBCollectionGetMetrics.json
 */
async function cosmosDbCollectionGetMetrics() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseRid = "databaseRid";
  const collectionRid = "collectionRid";
  const filter =
    "$filter=(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.collection.listMetrics(
    resourceGroupName,
    accountName,
    databaseRid,
    collectionRid,
    filter
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbCollectionGetMetrics().catch(console.error);
