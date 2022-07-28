import com.azure.core.util.Context;

/** Samples for CloudServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudService_List_ByResourceGroup.json
     */
    /**
     * Sample code: List Cloud Services in a Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServicesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .listByResourceGroup("ConstosoRG", Context.NONE);
    }
}
