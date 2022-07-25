const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete SQL Migration Service.
 *
 * @summary Delete SQL Migration Service.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/DeleteMigrationService.json
 */
async function deleteSqlMigrationService() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.beginDeleteAndWait(
    resourceGroupName,
    sqlMigrationServiceName
  );
  console.log(result);
}

deleteSqlMigrationService().catch(console.error);
