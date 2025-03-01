const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to All supported regulatory compliance controls details and state for selected standard
 *
 * @summary All supported regulatory compliance controls details and state for selected standard
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceControlList_example.json
 */
async function getAllRegulatoryComplianceControlsDetailsAndStateForSelectedStandard() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const regulatoryComplianceStandardName = "PCI-DSS-3.2";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.regulatoryComplianceControls.list(
    regulatoryComplianceStandardName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
