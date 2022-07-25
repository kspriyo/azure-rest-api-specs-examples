const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Dashboard.
 *
 * @summary Creates or updates a Dashboard.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/createOrUpdateDashboard.json
 */
async function createOrUpdateADashboard() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "testRG";
  const dashboardName = "testDashboard";
  const dashboard = {
    lenses: [
      {
        order: 1,
        parts: [
          { position: { colSpan: 3, rowSpan: 4, x: 1, y: 2 } },
          { position: { colSpan: 6, rowSpan: 6, x: 5, y: 5 } },
        ],
      },
      { order: 2, parts: [] },
    ],
    location: "eastus",
    metadata: { metadata: { ColSpan: 2, RowSpan: 1, X: 4, Y: 3 } },
    tags: { aKey: "aValue", anotherKey: "anotherValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const result = await client.dashboards.createOrUpdate(
    resourceGroupName,
    dashboardName,
    dashboard
  );
  console.log(result);
}

createOrUpdateADashboard().catch(console.error);
