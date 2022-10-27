import com.azure.core.util.Context;

/** Samples for BastionHosts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/BastionHostGet.json
     */
    /**
     * Sample code: Get Bastion Host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getBastionHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getBastionHosts()
            .getByResourceGroupWithResponse("rg1", "bastionhosttenant'", Context.NONE);
    }
}
