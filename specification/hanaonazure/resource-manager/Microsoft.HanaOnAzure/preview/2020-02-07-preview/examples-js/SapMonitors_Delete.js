const { HanaManagementClient } = require("@azure/arm-hanaonazure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a SAP monitor with the specified subscription, resource group, and monitor name.
 *
 * @summary Deletes a SAP monitor with the specified subscription, resource group, and monitor name.
 * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Delete.json
 */
async function deletesASapMonitor() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const sapMonitorName = "mySapMonitor";
  const credential = new DefaultAzureCredential();
  const client = new HanaManagementClient(credential, subscriptionId);
  const result = await client.sapMonitors.beginDeleteAndWait(resourceGroupName, sapMonitorName);
  console.log(result);
}

deletesASapMonitor().catch(console.error);
