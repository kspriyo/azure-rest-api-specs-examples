import com.azure.core.util.Context;

/** Samples for DatabaseBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseBlobAuditingGet.json
     */
    /**
     * Sample code: Get a database's blob auditing policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSBlobAuditingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseBlobAuditingPolicies()
            .getWithResponse("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
