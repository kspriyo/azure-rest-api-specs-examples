import com.azure.core.util.Context;

/** Samples for RestorableDroppedDatabases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/RestorableDroppedDatabaseList.json
     */
    /**
     * Sample code: Get list of restorable dropped databases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfRestorableDroppedDatabases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRestorableDroppedDatabases()
            .listByServer("restorabledroppeddatabasetest-1349", "restorabledroppeddatabasetest-1840", Context.NONE);
    }
}
