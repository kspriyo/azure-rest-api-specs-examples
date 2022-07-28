import com.azure.core.util.Context;

/** Samples for Databases ListMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseMetricsListWithFilter.json
     */
    /**
     * Sample code: List database usage metrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseUsageMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .listMetrics(
                "sqlcrudtest-6730",
                "sqlcrudtest-9007",
                "3481",
                "name/value eq 'cpu_percent' and timeGrain eq '00:10:00' and startTime eq '2017-06-02T18:35:00Z' and"
                    + " endTime eq '2017-06-02T18:55:00Z'",
                Context.NONE);
    }
}
