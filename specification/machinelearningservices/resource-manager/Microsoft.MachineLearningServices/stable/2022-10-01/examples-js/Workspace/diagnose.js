const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Diagnose workspace setup issue.
 *
 * @summary Diagnose workspace setup issue.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/diagnose.json
 */
async function diagnoseWorkspace() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-1234";
  const workspaceName = "testworkspace";
  const parameters = {
    value: {
      applicationInsights: {},
      containerRegistry: {},
      dnsResolution: {},
      keyVault: {},
      nsg: {},
      others: {},
      resourceLock: {},
      storageAccount: {},
      udr: {},
    },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaces.beginDiagnoseAndWait(
    resourceGroupName,
    workspaceName,
    options
  );
  console.log(result);
}

diagnoseWorkspace().catch(console.error);
