const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancel a remediation at resource scope.
 *
 * @summary Cancel a remediation at resource scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CancelResourceScope.json
 */
async function cancelARemediationAtIndividualResourceScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1";
  const remediationName = "myRemediation";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.remediations.cancelAtResource(resourceId, remediationName);
  console.log(result);
}

cancelARemediationAtIndividualResourceScope().catch(console.error);
