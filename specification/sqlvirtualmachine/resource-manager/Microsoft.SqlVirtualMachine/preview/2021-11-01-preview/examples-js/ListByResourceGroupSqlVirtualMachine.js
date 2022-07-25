const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all SQL virtual machines in a resource group.
 *
 * @summary Gets all SQL virtual machines in a resource group.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListByResourceGroupSqlVirtualMachine.json
 */
async function getsAllSqlVirtualMachinesInAResourceGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlVirtualMachines.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsAllSqlVirtualMachinesInAResourceGroup().catch(console.error);
