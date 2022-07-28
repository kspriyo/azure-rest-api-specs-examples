import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListSupported. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/ListSupportedPrivateLinkResources.json
     */
    /**
     * Sample code: ListSupportedPrivateLinkResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSupportedPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getPrivateLinkResources()
            .listSupported("rg1", "mysearchservice", null, Context.NONE);
    }
}
