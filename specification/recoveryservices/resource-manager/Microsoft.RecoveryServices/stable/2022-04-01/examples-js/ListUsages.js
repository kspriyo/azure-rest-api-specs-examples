const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches the usages of the vault.
 *
 * @summary Fetches the usages of the vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/ListUsages.json
 */
async function getsVaultUsages() {
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = "Default-RecoveryServices-ResourceGroup";
  const vaultName = "swaggerExample";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByVaults(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsVaultUsages().catch(console.error);
