import com.azure.core.util.Context;

/** Samples for SharedPrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/GetSharedPrivateLinkResource.json
     */
    /**
     * Sample code: SharedPrivateLinkResourceGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sharedPrivateLinkResourceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getSharedPrivateLinkResources()
            .getWithResponse("rg1", "mysearchservice", "testResource", null, Context.NONE);
    }
}
