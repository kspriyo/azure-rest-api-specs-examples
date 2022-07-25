const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The services resource is the top-level resource that represents the Database Migration Service. This action stops the service and the service cannot be used for data migration. The service owner won't be billed when the service is stopped.
 *
 * @summary The services resource is the top-level resource that represents the Database Migration Service. This action stops the service and the service cannot be used for data migration. The service owner won't be billed when the service is stopped.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Services_Stop.json
 */
async function servicesStop() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.services.beginStopAndWait(groupName, serviceName);
  console.log(result);
}

servicesStop().catch(console.error);
