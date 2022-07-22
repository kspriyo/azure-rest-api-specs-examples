const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB Gremlin graph from autoscale to manual throughput
 *
 * @summary Migrate an Azure Cosmos DB Gremlin graph from autoscale to manual throughput
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBGremlinGraphMigrateToManualThroughput.json
 */
async function cosmosDbGremlinGraphMigrateToManualThroughput() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const graphName = "graphName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginMigrateGremlinGraphToManualThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    graphName
  );
  console.log(result);
}

cosmosDbGremlinGraphMigrateToManualThroughput().catch(console.error);
