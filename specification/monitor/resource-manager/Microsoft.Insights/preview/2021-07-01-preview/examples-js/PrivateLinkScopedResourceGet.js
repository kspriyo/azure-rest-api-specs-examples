const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a scoped resource in a private link scope.
 *
 * @summary Gets a scoped resource in a private link scope.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceGet.json
 */
async function getsPrivateLinkScopedResource() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "MyResourceGroup";
  const scopeName = "MyPrivateLinkScope";
  const name = "scoped-resource-name";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.privateLinkScopedResources.get(resourceGroupName, scopeName, name);
  console.log(result);
}

getsPrivateLinkScopedResource().catch(console.error);
