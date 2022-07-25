const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the sensitivity label of a given column in a Sql pool
 *
 * @summary Creates or updates the sensitivity label of a given column in a Sql pool
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolColumnSensitivityLabelWithAllParameters.json
 */
async function updatesTheSensitivityLabelOfAGivenColumnWithAllParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const workspaceName = "myServer";
  const sqlPoolName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "myTable";
  const columnName = "myColumn";
  const parameters = {
    informationType: "PhoneNumber",
    informationTypeId: "d22fa6e9-5ee4-3bde-4c2b-a409604c4646",
    labelId: "bf91e08c-f4f0-478a-b016-25164b2a65ff",
    labelName: "PII",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolSensitivityLabels.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    schemaName,
    tableName,
    columnName,
    parameters
  );
  console.log(result);
}

updatesTheSensitivityLabelOfAGivenColumnWithAllParameters().catch(console.error);
