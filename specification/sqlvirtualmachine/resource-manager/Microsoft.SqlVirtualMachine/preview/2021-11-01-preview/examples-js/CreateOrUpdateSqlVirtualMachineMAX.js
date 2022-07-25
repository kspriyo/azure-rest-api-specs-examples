const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a SQL virtual machine.
 *
 * @summary Creates or updates a SQL virtual machine.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateSqlVirtualMachineMAX.json
 */
async function createsOrUpdatesASqlVirtualMachineWithMaxParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlVirtualMachineName = "testvm";
  const parameters = {
    assessmentSettings: {
      enable: true,
      runImmediately: true,
      schedule: {
        dayOfWeek: "Sunday",
        enable: true,
        monthlyOccurrence: undefined,
        startTime: "23:17",
        weeklyInterval: 1,
      },
    },
    autoBackupSettings: {
      backupScheduleType: "Manual",
      backupSystemDbs: true,
      enable: true,
      enableEncryption: true,
      fullBackupFrequency: "Daily",
      fullBackupStartTime: 6,
      fullBackupWindowHours: 11,
      logBackupFrequency: 10,
      password: "<Password>",
      retentionPeriod: 17,
      storageAccessKey: "<primary storage access key>",
      storageAccountUrl: "https://teststorage.blob.core.windows.net/",
      storageContainerName: "testcontainer",
    },
    autoPatchingSettings: {
      dayOfWeek: "Sunday",
      enable: true,
      maintenanceWindowDuration: 60,
      maintenanceWindowStartingHour: 2,
    },
    keyVaultCredentialSettings: { enable: false },
    location: "northeurope",
    serverConfigurationsManagementSettings: {
      additionalFeaturesServerConfigurations: { isRServicesEnabled: false },
      sqlConnectivityUpdateSettings: {
        connectivityType: "PRIVATE",
        port: 1433,
        sqlAuthUpdatePassword: "<password>",
        sqlAuthUpdateUserName: "sqllogin",
      },
      sqlInstanceSettings: {
        collation: "SQL_Latin1_General_CP1_CI_AS",
        isOptimizeForAdHocWorkloadsEnabled: true,
        maxDop: 8,
        maxServerMemoryMB: 128,
        minServerMemoryMB: 0,
      },
      sqlStorageUpdateSettings: {
        diskConfigurationType: "NEW",
        diskCount: 1,
        startingDeviceId: 2,
      },
      sqlWorkloadTypeUpdateSettings: { sqlWorkloadType: "OLTP" },
    },
    sqlImageSku: "Enterprise",
    sqlManagement: "Full",
    sqlServerLicenseType: "PAYG",
    virtualMachineResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    sqlVirtualMachineName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesASqlVirtualMachineWithMaxParameters().catch(console.error);
