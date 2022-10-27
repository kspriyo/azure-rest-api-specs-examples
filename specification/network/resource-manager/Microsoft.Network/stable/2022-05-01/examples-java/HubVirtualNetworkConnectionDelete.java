import com.azure.core.util.Context;

/** Samples for HubVirtualNetworkConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/HubVirtualNetworkConnectionDelete.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void hubVirtualNetworkConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getHubVirtualNetworkConnections()
            .delete("rg1", "virtualHub1", "connection1", Context.NONE);
    }
}
