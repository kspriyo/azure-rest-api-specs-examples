const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns an azure operation result.
 *
 * @summary Returns an azure operation result.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/OperationResults/getOperationResult.json
 */
async function getOperationResult() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const operationResultId = "a64149d8-84cb-4566-ab8e-b4ee1a074174";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.operationResults.get(operationResultId);
  console.log(result);
}

getOperationResult().catch(console.error);
