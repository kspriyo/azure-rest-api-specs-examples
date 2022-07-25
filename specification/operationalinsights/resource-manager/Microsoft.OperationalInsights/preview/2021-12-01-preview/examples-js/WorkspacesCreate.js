const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a workspace.
 *
 * @summary Create or update a workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesCreate.json
 */
async function workspacesCreate() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "oiautorest6685";
  const workspaceName = "oiautorest6685";
  const parameters = {
    location: "australiasoutheast",
    retentionInDays: 30,
    sku: { name: "PerGB2018" },
    tags: { tag1: "val1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.workspaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    parameters
  );
  console.log(result);
}

workspacesCreate().catch(console.error);
