const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List data versions in the data container
 *
 * @summary List data versions in the data container
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataVersionBase/list.json
 */
async function listDataVersionBase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const orderBy = "string";
  const top = 1;
  const tags = "string";
  const options = { orderBy, top, tags };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataVersions.list(
    resourceGroupName,
    workspaceName,
    name,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDataVersionBase().catch(console.error);
