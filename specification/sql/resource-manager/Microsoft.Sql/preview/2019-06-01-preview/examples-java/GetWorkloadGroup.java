import com.azure.core.util.Context;

/** Samples for WorkloadGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/GetWorkloadGroup.json
     */
    /**
     * Sample code: Gets a workload group for a data warehouse.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAWorkloadGroupForADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getWorkloadGroups()
            .getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", Context.NONE);
    }
}
