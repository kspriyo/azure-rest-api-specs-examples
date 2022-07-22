const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB Gremlin graph
 *
 * @summary Create or update an Azure Cosmos DB Gremlin graph
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBGremlinGraphCreateUpdate.json
 */
async function cosmosDbGremlinGraphCreateUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const graphName = "graphName";
  const createUpdateGremlinGraphParameters = {
    location: "West US",
    options: {},
    resource: {
      conflictResolutionPolicy: {
        conflictResolutionPath: "/path",
        mode: "LastWriterWins",
      },
      defaultTtl: 100,
      id: "graphName",
      indexingPolicy: {
        automatic: true,
        excludedPaths: [],
        includedPaths: [
          {
            path: "/*",
            indexes: [
              { dataType: "String", kind: "Range", precision: -1 },
              { dataType: "Number", kind: "Range", precision: -1 },
            ],
          },
        ],
        indexingMode: "consistent",
      },
      partitionKey: { kind: "Hash", paths: ["/AccountNumber"] },
      uniqueKeyPolicy: { uniqueKeys: [{ paths: ["/testPath"] }] },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginCreateUpdateGremlinGraphAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    graphName,
    createUpdateGremlinGraphParameters
  );
  console.log(result);
}

cosmosDbGremlinGraphCreateUpdate().catch(console.error);
