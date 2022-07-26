const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a database
 *
 * @summary Creates a database
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesCreate.json
 */
async function redisEnterpriseDatabasesCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = {
    clientProtocol: "Encrypted",
    clusteringPolicy: "EnterpriseCluster",
    evictionPolicy: "AllKeysLRU",
    modules: [
      { name: "RedisBloom", args: "ERROR_RATE 0.00 INITIAL_SIZE 400" },
      { name: "RedisTimeSeries", args: "RETENTION_POLICY 20" },
      { name: "RediSearch" },
    ],
    persistence: { aofEnabled: true, aofFrequency: "1s" },
    port: 10000,
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters
  );
  console.log(result);
}

redisEnterpriseDatabasesCreate().catch(console.error);
