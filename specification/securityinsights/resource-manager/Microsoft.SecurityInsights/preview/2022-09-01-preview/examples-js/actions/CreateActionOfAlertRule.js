const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the action of alert rule.
 *
 * @summary Creates or updates the action of alert rule.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/actions/CreateActionOfAlertRule.json
 */
async function createsOrUpdatesAnActionOfAlertRule() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const ruleId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
  const actionId = "912bec42-cb66-4c03-ac63-1761b6898c3e";
  const action = {
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    logicAppResourceId:
      "/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/MyAlerts",
    triggerUri:
      "https://prod-31.northcentralus.logic.azure.com:443/workflows/cd3765391efd48549fd7681ded1d48d7/triggers/manual/paths/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=signature",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.actions.createOrUpdate(
    resourceGroupName,
    workspaceName,
    ruleId,
    actionId,
    action
  );
  console.log(result);
}

createsOrUpdatesAnActionOfAlertRule().catch(console.error);
