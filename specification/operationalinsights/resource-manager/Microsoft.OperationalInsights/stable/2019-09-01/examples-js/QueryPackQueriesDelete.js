const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specific Query defined within an Log Analytics QueryPack.
 *
 * @summary Deletes a specific Query defined within an Log Analytics QueryPack.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesDelete.json
 */
async function queryDelete() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
  const resourceGroupName = "my-resource-group";
  const queryPackName = "my-querypack";
  const id = "a449f8af-8e64-4b3a-9b16-5a7165ff98c4";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.queries.delete(resourceGroupName, queryPackName, id);
  console.log(result);
}

queryDelete().catch(console.error);
