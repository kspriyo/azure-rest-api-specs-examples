const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to update existing IoT Security solution tags or user defined resources. To update other fields use the CreateOrUpdate method.
 *
 * @summary Use this method to update existing IoT Security solution tags or user defined resources. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
 */
async function useThisMethodToUpdateExistingIoTSecuritySolution() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const solutionName = "default";
  const updateIotSecuritySolutionData = {
    recommendationsConfiguration: [
      { recommendationType: "IoT_OpenPorts", status: "Disabled" },
      { recommendationType: "IoT_SharedCredentials", status: "Disabled" },
    ],
    tags: { foo: "bar" },
    userDefinedResources: {
      query: 'where type != "microsoft.devices/iothubs" | where name contains "v2"',
      querySubscriptions: ["075423e9-7d33-4166-8bdf-3920b04e3735"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.iotSecuritySolution.update(
    resourceGroupName,
    solutionName,
    updateIotSecuritySolutionData
  );
  console.log(result);
}
