const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The project resource is a nested resource representing a stored migration project. This method returns a list of projects owned by a service resource.
 *
 * @summary The project resource is a nested resource representing a stored migration project. This method returns a list of projects owned by a service resource.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Projects_List.json
 */
async function projectsList() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projects.list(groupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

projectsList().catch(console.error);
