import com.azure.core.util.Context;

/** Samples for ElasticPoolDatabaseActivities ListByElasticPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ElasticPoolDatabaseActivityList.json
     */
    /**
     * Sample code: List elastic pool database activity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listElasticPoolDatabaseActivity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPoolDatabaseActivities()
            .listByElasticPool("sqlcrudtest-4673", "sqlcrudtest-603", "7537", Context.NONE);
    }
}
