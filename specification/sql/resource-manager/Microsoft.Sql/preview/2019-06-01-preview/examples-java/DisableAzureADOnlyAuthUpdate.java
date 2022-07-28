import com.azure.core.util.Context;

/** Samples for ServerAzureADAdministrators DisableAzureADOnlyAuthentication. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/DisableAzureADOnlyAuthUpdate.json
     */
    /**
     * Sample code: Disables Azure Active Directory only authentication on logical server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void disablesAzureActiveDirectoryOnlyAuthenticationOnLogicalServer(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerAzureADAdministrators()
            .disableAzureADOnlyAuthentication("sqlcrudtest-4799", "sqlcrudtest-6440", Context.NONE);
    }
}
