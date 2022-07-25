const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Data Lake Analytics compute policies within the specified Data Lake Analytics account. An account supports, at most, 50 policies
 *
 * @summary Lists the Data Lake Analytics compute policies within the specified Data Lake Analytics account. An account supports, at most, 50 policies
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_ListByAccount.json
 */
async function listsTheComputePoliciesWithinTheAdlaAccount() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.computePolicies.listByAccount(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsTheComputePoliciesWithinTheAdlaAccount().catch(console.error);
