const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a data connection.
 *
 * @summary Creates or updates a data connection.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCreateOrUpdate.json
 */
async function kustoPoolDataConnectionsCreateOrUpdateJson() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const dataConnectionName = "DataConnections8";
  const parameters = {
    consumerGroup: "testConsumerGroup1",
    eventHubResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1",
    kind: "EventHub",
    location: "westus",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDataConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    dataConnectionName,
    parameters
  );
  console.log(result);
}

kustoPoolDataConnectionsCreateOrUpdateJson().catch(console.error);
