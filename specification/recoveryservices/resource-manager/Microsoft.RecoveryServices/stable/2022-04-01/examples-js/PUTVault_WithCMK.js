const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Recovery Services vault.
 *
 * @summary Creates or updates a Recovery Services vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/PUTVault_WithCMK.json
 */
async function createOrUpdateVaultWithCustomerManagedKeys() {
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = "Default-RecoveryServices-ResourceGroup";
  const vaultName = "swaggerExample";
  const vault = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/85bf5e8c30844f42Add2746ebb7e97b2/resourcegroups/defaultrg/providers/MicrosoftManagedIdentity/userAssignedIdentities/examplemsi":
          {},
      },
    },
    location: "West US",
    properties: {
      encryption: {
        infrastructureEncryption: "Enabled",
        kekIdentity: {
          userAssignedIdentity:
            "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
        },
        keyVaultProperties: {
          keyUri: "https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3",
        },
      },
    },
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vaultName,
    vault
  );
  console.log(result);
}

createOrUpdateVaultWithCustomerManagedKeys().catch(console.error);
