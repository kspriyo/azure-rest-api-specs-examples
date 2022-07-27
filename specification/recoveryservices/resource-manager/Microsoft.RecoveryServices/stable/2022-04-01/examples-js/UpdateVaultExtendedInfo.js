const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create vault extended info.
 *
 * @summary Create vault extended info.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/UpdateVaultExtendedInfo.json
 */
async function putExtendedInfoOfResource() {
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = "Default-RecoveryServices-ResourceGroup";
  const vaultName = "swaggerExample";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaultExtendedInfo.createOrUpdate(resourceGroupName, vaultName, {
    integrityKey: "J99wzS27fmJ+Wjot7xO5wA==",
    algorithm: "None",
  });
  console.log(result);
}

putExtendedInfoOfResource().catch(console.error);
