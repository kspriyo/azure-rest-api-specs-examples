const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details and state of assessments mapped to selected regulatory compliance control
 *
 * @summary Details and state of assessments mapped to selected regulatory compliance control
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceAssessmentList_example.json
 */
async function getAllAssessmentsMappedToSelectedRegulatoryComplianceControl() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const regulatoryComplianceStandardName = "PCI-DSS-3.2";
  const regulatoryComplianceControlName = "1.1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.regulatoryComplianceAssessments.list(
    regulatoryComplianceStandardName,
    regulatoryComplianceControlName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
