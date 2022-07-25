const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the sensitivity label of a given column in a Sql pool
 *
 * @summary Deletes the sensitivity label of a given column in a Sql pool
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolColumnSensitivityLabel.json
 */
async function deletesTheSensitivityLabelOfAGivenColumn() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const workspaceName = "myServer";
  const sqlPoolName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "myTable";
  const columnName = "myColumn";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolSensitivityLabels.delete(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    schemaName,
    tableName,
    columnName
  );
  console.log(result);
}

deletesTheSensitivityLabelOfAGivenColumn().catch(console.error);
