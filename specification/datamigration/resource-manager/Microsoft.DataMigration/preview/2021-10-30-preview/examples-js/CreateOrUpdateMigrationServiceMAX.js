const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update SQL Migration Service.
 *
 * @summary Create or Update SQL Migration Service.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/CreateOrUpdateMigrationServiceMAX.json
 */
async function createOrUpdateSqlMigrationServiceWithMaximumParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "testagent";
  const parameters = { location: "northeurope" };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    sqlMigrationServiceName,
    parameters
  );
  console.log(result);
}

createOrUpdateSqlMigrationServiceWithMaximumParameters().catch(console.error);
