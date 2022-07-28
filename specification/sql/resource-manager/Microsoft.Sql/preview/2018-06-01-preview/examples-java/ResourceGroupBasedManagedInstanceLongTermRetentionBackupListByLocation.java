import com.azure.core.util.Context;

/** Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByLocation.json
     */
    /**
     * Sample code: Get all long term retention backups under the location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllLongTermRetentionBackupsUnderTheLocation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getLongTermRetentionManagedInstanceBackups()
            .listByResourceGroupLocation("testResourceGroup", "japaneast", null, null, Context.NONE);
    }
}
