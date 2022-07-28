import com.azure.core.util.Context;

/** Samples for FailoverGroups Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/FailoverGroupFailover.json
     */
    /**
     * Sample code: Planned failover of a failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void plannedFailoverOfAFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getFailoverGroups()
            .failover("Default", "failover-group-secondary-server", "failover-group-test-3", Context.NONE);
    }
}
