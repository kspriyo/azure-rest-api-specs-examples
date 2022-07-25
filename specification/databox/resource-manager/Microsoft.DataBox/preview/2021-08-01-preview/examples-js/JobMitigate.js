const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Request to mitigate for a given job
 *
 * @summary Request to mitigate for a given job
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobMitigate.json
 */
async function mitigate() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const jobName = "SdkJob8367";
  const resourceGroupName = "SdkRg9836";
  const mitigateJobRequest = {
    customerResolutionCode: "MoveToCleanUpDevice",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.mitigate(jobName, resourceGroupName, mitigateJobRequest);
  console.log(result);
}

mitigate().catch(console.error);
