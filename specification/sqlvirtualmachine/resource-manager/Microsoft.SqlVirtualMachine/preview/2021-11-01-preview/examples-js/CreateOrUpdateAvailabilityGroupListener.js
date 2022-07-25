const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an availability group listener.
 *
 * @summary Creates or updates an availability group listener.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
 */
async function createsOrUpdatesAnAvailabilityGroupListener() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlVirtualMachineGroupName = "testvmgroup";
  const availabilityGroupListenerName = "agl-test";
  const parameters = {
    availabilityGroupName: "ag-test",
    loadBalancerConfigurations: [
      {
        loadBalancerResourceId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test",
        privateIpAddress: {
          ipAddress: "10.1.0.112",
          subnetResourceId:
            "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
        },
        probePort: 59983,
        sqlVirtualMachineInstances: [
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2",
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3",
        ],
      },
    ],
    port: 1433,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.availabilityGroupListeners.beginCreateOrUpdateAndWait(
    resourceGroupName,
    sqlVirtualMachineGroupName,
    availabilityGroupListenerName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesAnAvailabilityGroupListener().catch(console.error);
