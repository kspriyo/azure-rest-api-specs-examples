const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the cluster customer user credentials for the dedicated appliance.
 *
 * @summary Returns the cluster customer user credentials for the dedicated appliance.
 * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListClusterCustomerUserCredential.json
 */
async function listClusterCustomerUserCredentialAppliance() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "testresourcegroup";
  const resourceName = "appliance01";
  const credential = new DefaultAzureCredential();
  const client = new ResourceConnectorManagementClient(credential, subscriptionId);
  const result = await client.appliances.listClusterCustomerUserCredential(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

listClusterCustomerUserCredentialAppliance().catch(console.error);
