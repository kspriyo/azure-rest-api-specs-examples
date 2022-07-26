const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an existing (overwrite/recreate, with potential downtime) cache cluster
 *
 * @summary Creates or updates an existing (overwrite/recreate, with potential downtime) cache cluster
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseCreate.json
 */
async function redisEnterpriseCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const parameters = {
    location: "West US",
    minimumTlsVersion: "1.2",
    sku: { name: "EnterpriseFlash_F300", capacity: 3 },
    tags: { tag1: "value1" },
    zones: ["1", "2", "3"],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.redisEnterprise.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

redisEnterpriseCreate().catch(console.error);
