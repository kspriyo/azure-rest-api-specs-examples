import com.azure.core.util.Context;

/** Samples for ExtendedDatabaseBlobAuditingPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseExtendedAuditingSettingsList.json
     */
    /**
     * Sample code: List extended auditing settings of a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExtendedAuditingSettingsOfADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getExtendedDatabaseBlobAuditingPolicies()
            .listByDatabase("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
