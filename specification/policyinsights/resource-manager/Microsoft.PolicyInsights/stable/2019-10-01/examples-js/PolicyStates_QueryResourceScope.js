const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the resource.
 *
 * @summary Queries policy states for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryResourceScope.json
 */
async function queryAllPolicyStatesAtResourceScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "default";
  const resourceId =
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForResource(
    policyStatesResource,
    resourceId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryAllPolicyStatesAtResourceScope().catch(console.error);
