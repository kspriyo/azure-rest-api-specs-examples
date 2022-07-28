import com.azure.core.util.Context;

/** Samples for ManagedInstanceOperations ListByManagedInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ListManagedInstanceOperations.json
     */
    /**
     * Sample code: List the managed instance management operations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheManagedInstanceManagementOperations(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceOperations()
            .listByManagedInstance("sqlcrudtest-7398", "sqlcrudtest-4645", Context.NONE);
    }
}
