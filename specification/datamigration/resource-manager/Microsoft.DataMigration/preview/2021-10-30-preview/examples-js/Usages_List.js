const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This method returns region-specific quotas and resource usage information for the Database Migration Service.
 *
 * @summary This method returns region-specific quotas and resource usage information for the Database Migration Service.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Usages_List.json
 */
async function servicesUsages() {
  const subscriptionId = "90fb80a6-0f71-4761-8f03-921e7396f3c0";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

servicesUsages().catch(console.error);
