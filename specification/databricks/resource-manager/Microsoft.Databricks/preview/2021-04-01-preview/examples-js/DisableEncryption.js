const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new workspace.
 *
 * @summary Creates a new workspace.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/DisableEncryption.json
 */
async function revertCustomerManagedKeyCmkEncryptionToMicrosoftManagedKeysEncryptionOnAWorkspace() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const workspaceName = "myWorkspace";
  const parameters = {
    location: "westus",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
    parameters: { encryption: { value: { keySource: "Default" } } },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    parameters
  );
  console.log(result);
}

revertCustomerManagedKeyCmkEncryptionToMicrosoftManagedKeysEncryptionOnAWorkspace().catch(
  console.error
);
