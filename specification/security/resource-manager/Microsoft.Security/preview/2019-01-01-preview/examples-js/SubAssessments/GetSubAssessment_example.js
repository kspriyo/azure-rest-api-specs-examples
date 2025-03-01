const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a security sub-assessment on your scanned resource
 *
 * @summary Get a security sub-assessment on your scanned resource
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/GetSubAssessment_example.json
 */
async function getSecurityRecommendationTaskFromSecurityDataLocation() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope =
    "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2";
  const assessmentName = "1195afff-c881-495e-9bc5-1486211ae03f";
  const subAssessmentName = "95f7da9c-a2a4-1322-0758-fcd24ef09b85";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.subAssessments.get(scope, assessmentName, subAssessmentName);
  console.log(result);
}
