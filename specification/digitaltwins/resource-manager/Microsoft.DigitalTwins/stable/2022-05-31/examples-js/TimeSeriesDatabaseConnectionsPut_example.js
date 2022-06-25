const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a time series database connection.
 *
 * @summary Create or update a time series database connection.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/TimeSeriesDatabaseConnectionsPut_example.json
 */
async function createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstance() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const timeSeriesDatabaseConnectionName = "myConnection";
  const timeSeriesDatabaseConnectionDescription = {
    properties: {
      adxDatabaseName: "myDatabase",
      adxEndpointUri: "https://mycluster.kusto.windows.net",
      adxResourceId:
        "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster",
      adxTableName: "myTable",
      connectionType: "AzureDataExplorer",
      eventHubEndpointUri: "sb://myeh.servicebus.windows.net/",
      eventHubEntityPath: "myeh",
      eventHubNamespaceResourceId:
        "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.timeSeriesDatabaseConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    timeSeriesDatabaseConnectionName,
    timeSeriesDatabaseConnectionDescription
  );
  console.log(result);
}

createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstance().catch(console.error);
