const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint connection associated with the redis cache.
 *
 * @summary Deletes the specified private endpoint connection associated with the redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheDeletePrivateEndpointConnection.json
 */
async function redisCacheDeletePrivateEndpointConnection() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rgtest01";
  const cacheName = "cachetest01";
  const privateEndpointConnectionName = "pectest01";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    cacheName,
    privateEndpointConnectionName
  );
  console.log(result);
}

redisCacheDeletePrivateEndpointConnection().catch(console.error);
