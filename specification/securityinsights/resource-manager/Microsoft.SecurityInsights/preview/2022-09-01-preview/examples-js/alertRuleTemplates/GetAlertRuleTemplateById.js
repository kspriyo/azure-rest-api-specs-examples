const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the alert rule template.
 *
 * @summary Gets the alert rule template.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplateById.json
 */
async function getAlertRuleTemplateById() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const alertRuleTemplateId = "65360bb0-8986-4ade-a89d-af3cf44d28aa";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRuleTemplates.get(
    resourceGroupName,
    workspaceName,
    alertRuleTemplateId
  );
  console.log(result);
}

getAlertRuleTemplateById().catch(console.error);
