const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Supported regulatory compliance details and state for selected assessment
 *
 * @summary Supported regulatory compliance details and state for selected assessment
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceAssessment_example.json
 */
async function getSelectedRegulatoryComplianceAssessmentDetailsAndState() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const regulatoryComplianceStandardName = "PCI-DSS-3.2";
  const regulatoryComplianceControlName = "1.1";
  const regulatoryComplianceAssessmentName = "968548cb-02b3-8cd2-11f8-0cf64ab1a347";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.regulatoryComplianceAssessments.get(
    regulatoryComplianceStandardName,
    regulatoryComplianceControlName,
    regulatoryComplianceAssessmentName
  );
  console.log(result);
}
