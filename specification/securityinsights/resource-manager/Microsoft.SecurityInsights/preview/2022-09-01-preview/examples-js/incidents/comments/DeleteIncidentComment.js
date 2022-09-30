const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the incident comment.
 *
 * @summary Delete the incident comment.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/comments/DeleteIncidentComment.json
 */
async function deleteTheIncidentComment() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const incidentId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
  const incidentCommentId = "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.incidentComments.delete(
    resourceGroupName,
    workspaceName,
    incidentId,
    incidentCommentId
  );
  console.log(result);
}

deleteTheIncidentComment().catch(console.error);
