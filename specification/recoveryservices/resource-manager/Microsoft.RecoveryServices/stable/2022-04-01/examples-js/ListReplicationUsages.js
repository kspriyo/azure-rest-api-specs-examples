const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches the replication usages of the vault.
 *
 * @summary Fetches the replication usages of the vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/ListReplicationUsages.json
 */
async function getsReplicationUsagesOfVault() {
  const subscriptionId = "6808dbbc-98c7-431f-a1b1-9580902423b7";
  const resourceGroupName = "avrai7517RG1";
  const vaultName = "avrai7517Vault1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationUsages.list(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsReplicationUsagesOfVault().catch(console.error);
