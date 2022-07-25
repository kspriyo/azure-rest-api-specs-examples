const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update Database Migration resource.
 *
 * @summary Create or Update Database Migration resource.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/SqlMiCreateOrUpdateDatabaseMigrationMAX.json
 */
async function createOrUpdateDatabaseMigrationResourceWithMaximumParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const managedInstanceName = "managedInstance1";
  const targetDbName = "db1";
  const parameters = {
    properties: {
      backupConfiguration: {
        sourceLocation: {
          fileShare: {
            path: "C:aaa\bbbccc",
            password: "placeholder",
            username: "name",
          },
        },
        targetLocation: {
          accountKey: "abcd",
          storageAccountResourceId: "account.database.windows.net",
        },
      },
      kind: "SqlMi",
      migrationService:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent",
      offlineConfiguration: {
        lastBackupName: "last_backup_file_name",
        offline: true,
      },
      scope:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance",
      sourceDatabaseName: "aaa",
      sourceSqlConnection: {
        authentication: "WindowsAuthentication",
        dataSource: "aaa",
        encryptConnection: true,
        password: "placeholder",
        trustServerCertificate: true,
        userName: "bbb",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlMi.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    targetDbName,
    parameters
  );
  console.log(result);
}

createOrUpdateDatabaseMigrationResourceWithMaximumParameters().catch(console.error);
