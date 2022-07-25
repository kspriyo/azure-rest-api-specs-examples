import com.azure.core.util.Context;

/** Samples for SimGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupListByResourceGroup.json
     */
    /**
     * Sample code: List SIM groups in a resource group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMGroupsInAResourceGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().listByResourceGroup("rg1", Context.NONE);
    }
}
