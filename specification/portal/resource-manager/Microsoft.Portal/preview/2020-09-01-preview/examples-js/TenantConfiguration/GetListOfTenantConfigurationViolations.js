const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of items that violate tenant's configuration.
 *
 * @summary Gets list of items that violate tenant's configuration.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/GetListOfTenantConfigurationViolations.json
 */
async function getListOfItemsThatViolateTenantConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listTenantConfigurationViolations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfItemsThatViolateTenantConfiguration().catch(console.error);
