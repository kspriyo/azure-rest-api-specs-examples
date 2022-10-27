import com.azure.core.util.Context;

/** Samples for AvailableServiceAliases List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AvailableServiceAliasesList.json
     */
    /**
     * Sample code: Get available service aliases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableServiceAliases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailableServiceAliases().list("westcentralus", Context.NONE);
    }
}
