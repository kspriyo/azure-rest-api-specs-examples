const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of private link resources that need to be created for Backup and SiteRecovery
 *
 * @summary Returns the list of private link resources that need to be created for Backup and SiteRecovery
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/ListPrivateLinkResources.json
 */
async function listPrivateLinkResources() {
  const subscriptionId = "6c48fa17-39c7-45f1-90ac-47a587128ace";
  const resourceGroupName = "petesting";
  const vaultName = "pemsi-ecy-rsv2";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResourcesOperations.list(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPrivateLinkResources().catch(console.error);
