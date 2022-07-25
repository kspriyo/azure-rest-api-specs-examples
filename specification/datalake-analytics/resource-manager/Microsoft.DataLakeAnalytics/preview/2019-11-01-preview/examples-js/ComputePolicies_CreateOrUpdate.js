const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified compute policy. During update, the compute policy with the specified name will be replaced with this new compute policy. An account supports, at most, 50 policies
 *
 * @summary Creates or updates the specified compute policy. During update, the compute policy with the specified name will be replaced with this new compute policy. An account supports, at most, 50 policies
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_CreateOrUpdate.json
 */
async function createsOrUpdatesTheSpecifiedComputePolicy() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const computePolicyName = "test_policy";
  const parameters = {
    maxDegreeOfParallelismPerJob: 10,
    minPriorityPerJob: 30,
    objectId: "776b9091-8916-4638-87f7-9c989a38da98",
    objectType: "User",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.computePolicies.createOrUpdate(
    resourceGroupName,
    accountName,
    computePolicyName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesTheSpecifiedComputePolicy().catch(console.error);
