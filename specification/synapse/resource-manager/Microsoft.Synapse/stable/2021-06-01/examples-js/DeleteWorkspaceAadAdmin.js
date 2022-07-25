const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a workspace active directory admin
 *
 * @summary Deletes a workspace active directory admin
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspaceAadAdmin.json
 */
async function deleteWorkspaceActiveDirectoryAdmin() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceAadAdmins.beginDeleteAndWait(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

deleteWorkspaceActiveDirectoryAdmin().catch(console.error);
