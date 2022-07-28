import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/** Samples for ServerBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ServerBlobAuditingCreateMax.json
     */
    /**
     * Sample code: Update a server's blob auditing policy with all parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServerSBlobAuditingPolicyWithAllParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerBlobAuditingPolicies()
            .createOrUpdate(
                "blobauditingtest-4799",
                "blobauditingtest-6440",
                new ServerBlobAuditingPolicyInner()
                    .withState(BlobAuditingPolicyState.ENABLED)
                    .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey(
                        "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
                    .withRetentionDays(6)
                    .withAuditActionsAndGroups(
                        Arrays
                            .asList(
                                "SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP",
                                "FAILED_DATABASE_AUTHENTICATION_GROUP",
                                "BATCH_COMPLETED_GROUP"))
                    .withStorageAccountSubscriptionId(UUID.fromString("00000000-1234-0000-5678-000000000000"))
                    .withIsStorageSecondaryKeyInUse(false)
                    .withIsAzureMonitorTargetEnabled(true)
                    .withQueueDelayMs(4000),
                Context.NONE);
    }
}
