const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete container group virtual network association links. The operation does not delete other resources provided by the user.
 *
 * @summary Delete container group virtual network association links. The operation does not delete other resources provided by the user.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/SubnetServiceAssociationLinkDelete.json
 */
async function subnetServiceAssociationLinkDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const virtualNetworkName = "demo2";
  const subnetName = "demo3";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.subnetServiceAssociationLink.beginDeleteAndWait(
    resourceGroupName,
    virtualNetworkName,
    subnetName
  );
  console.log(result);
}

subnetServiceAssociationLinkDelete().catch(console.error);
