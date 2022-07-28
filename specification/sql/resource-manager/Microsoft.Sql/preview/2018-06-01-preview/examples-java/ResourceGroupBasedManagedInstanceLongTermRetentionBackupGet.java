import com.azure.core.util.Context;

/** Samples for LongTermRetentionManagedInstanceBackups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupGet.json
     */
    /**
     * Sample code: Get the long term retention backup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheLongTermRetentionBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getLongTermRetentionManagedInstanceBackups()
            .getByResourceGroupWithResponse(
                "testResourceGroup",
                "japaneast",
                "testInstance",
                "testDatabase",
                "55555555-6666-7777-8888-999999999999;131637960820000000",
                Context.NONE);
    }
}
