import com.azure.core.util.Context;

/** Samples for ServerKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the server key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteTheServerKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerKeys()
            .delete(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                "someVault_someKey_01234567890123456789012345678901",
                Context.NONE);
    }
}
