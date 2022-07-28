import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for SQL.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateLinkResourcesForSQL(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getPrivateLinkResources()
            .listByServer("Default", "test-svr", Context.NONE);
    }
}
