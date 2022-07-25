const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add a list of language extensions that can run within KQL queries.
 *
 * @summary Add a list of language extensions that can run within KQL queries.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsAdd.json
 */
async function kustoPoolAddLanguageExtensions() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = "kustorptest";
  const languageExtensionsToAdd = {
    value: [{ languageExtensionName: "PYTHON" }, { languageExtensionName: "R" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.beginAddLanguageExtensionsAndWait(
    workspaceName,
    kustoPoolName,
    resourceGroupName,
    languageExtensionsToAdd
  );
  console.log(result);
}

kustoPoolAddLanguageExtensions().catch(console.error);
