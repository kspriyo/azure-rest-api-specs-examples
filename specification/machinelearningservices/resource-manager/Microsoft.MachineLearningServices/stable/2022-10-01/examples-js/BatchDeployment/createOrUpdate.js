const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates/updates a batch inference deployment (asynchronous).
 *
 * @summary Creates/updates a batch inference deployment (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchDeployment/createOrUpdate.json
 */
async function createOrUpdateBatchDeployment() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const deploymentName = "testDeploymentName";
  const body = {
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: { string: {} },
    },
    kind: "string",
    location: "string",
    properties: {
      description: "string",
      codeConfiguration: { codeId: "string", scoringScript: "string" },
      compute: "string",
      environmentId: "string",
      environmentVariables: { string: "string" },
      errorThreshold: 1,
      loggingLevel: "Info",
      maxConcurrencyPerInstance: 1,
      miniBatchSize: 1,
      model: { assetId: "string", referenceType: "Id" },
      outputAction: "SummaryOnly",
      outputFileName: "string",
      properties: { string: "string" },
      resources: {
        instanceCount: 1,
        instanceType: "string",
        properties: { string: { "cd3c37dc-2876-4ca4-8a54-21bd7619724a": null } },
      },
      retrySettings: { maxRetries: 1, timeout: "PT5M" },
    },
    sku: {
      name: "string",
      capacity: 1,
      family: "string",
      size: "string",
      tier: "Free",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.batchDeployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName,
    body
  );
  console.log(result);
}

createOrUpdateBatchDeployment().catch(console.error);
