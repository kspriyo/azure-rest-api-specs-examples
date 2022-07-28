import com.azure.core.util.Context;

/** Samples for ServerAzureADAdministrators ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/AdministratorList.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory administrator.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfAzureActiveDirectoryAdministrator(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerAzureADAdministrators()
            .listByServer("sqlcrudtest-4799", "sqlcrudtest-6440", Context.NONE);
    }
}
