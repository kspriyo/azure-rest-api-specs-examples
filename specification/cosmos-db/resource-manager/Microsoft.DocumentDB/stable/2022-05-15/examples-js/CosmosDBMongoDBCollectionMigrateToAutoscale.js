const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB MongoDB collection from manual throughput to autoscale
 *
 * @summary Migrate an Azure Cosmos DB MongoDB collection from manual throughput to autoscale
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBMongoDBCollectionMigrateToAutoscale.json
 */
async function cosmosDbMongoDbcollectionMigrateToAutoscale() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const collectionName = "collectionName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginMigrateMongoDBCollectionToAutoscaleAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    collectionName
  );
  console.log(result);
}

cosmosDbMongoDbcollectionMigrateToAutoscale().catch(console.error);
