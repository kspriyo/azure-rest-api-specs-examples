const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the description of an existing time series database connection.
 *
 * @summary Get the description of an existing time series database connection.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/TimeSeriesDatabaseConnectionsGet_example.json
 */
async function getTimeSeriesDatabaseConnectionForADigitalTwinsInstance() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const timeSeriesDatabaseConnectionName = "myConnection";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.timeSeriesDatabaseConnections.get(
    resourceGroupName,
    resourceName,
    timeSeriesDatabaseConnectionName
  );
  console.log(result);
}

getTimeSeriesDatabaseConnectionForADigitalTwinsInstance().catch(console.error);
