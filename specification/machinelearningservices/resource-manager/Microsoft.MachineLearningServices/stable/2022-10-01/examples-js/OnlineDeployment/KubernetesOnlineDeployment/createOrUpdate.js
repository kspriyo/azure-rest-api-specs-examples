const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update Inference Endpoint Deployment (asynchronous).
 *
 * @summary Create or update Inference Endpoint Deployment (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/KubernetesOnlineDeployment/createOrUpdate.json
 */
async function createOrUpdateKubernetesOnlineDeployment() {
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
      appInsightsEnabled: false,
      codeConfiguration: { codeId: "string", scoringScript: "string" },
      containerResourceRequirements: {
        containerResourceLimits: { cpu: '"1"', gpu: '"1"', memory: '"2Gi"' },
        containerResourceRequests: { cpu: '"1"', gpu: '"1"', memory: '"2Gi"' },
      },
      endpointComputeType: "Kubernetes",
      environmentId: "string",
      environmentVariables: { string: "string" },
      instanceType: "string",
      livenessProbe: {
        failureThreshold: 1,
        initialDelay: "PT5M",
        period: "PT5M",
        successThreshold: 1,
        timeout: "PT5M",
      },
      model: "string",
      modelMountPath: "string",
      properties: { string: "string" },
      requestSettings: {
        maxConcurrentRequestsPerInstance: 1,
        maxQueueWait: "PT5M",
        requestTimeout: "PT5M",
      },
      scaleSettings: { scaleType: "Default" },
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
  const result = await client.onlineDeployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName,
    body
  );
  console.log(result);
}

createOrUpdateKubernetesOnlineDeployment().catch(console.error);
