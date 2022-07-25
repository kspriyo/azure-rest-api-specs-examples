const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the managed HSM Pools associated with the subscription.
 *
 * @summary The List operation gets information about the managed HSM Pools associated with the subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/ManagedHsm_ListBySubscription.json
 */
async function listManagedHsmPoolsInASubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedHsms.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedHsmPoolsInASubscription().catch(console.error);
