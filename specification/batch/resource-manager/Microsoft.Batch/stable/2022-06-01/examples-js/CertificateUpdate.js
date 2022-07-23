const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing certificate.
 *
 * @summary Updates the properties of an existing certificate.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-06-01/examples/CertificateUpdate.json
 */
async function updateCertificate() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const certificateName = "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e";
  const parameters = {
    data: "MIIJsgIBAzCCCW4GCSqGSIb3DQE...",
    password: "<ExamplePassword>",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.certificateOperations.update(
    resourceGroupName,
    accountName,
    certificateName,
    parameters
  );
  console.log(result);
}

updateCertificate().catch(console.error);
