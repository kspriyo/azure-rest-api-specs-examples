import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/** Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/UpdateManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Update the short term retention policy for the restorable dropped database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheShortTermRetentionPolicyForTheRestorableDroppedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies()
            .update(
                "resourceGroup",
                "testsvr",
                "testdb,131403269876900000",
                ManagedShortTermRetentionPolicyName.DEFAULT,
                new ManagedBackupShortTermRetentionPolicyInner(),
                Context.NONE);
    }
}
