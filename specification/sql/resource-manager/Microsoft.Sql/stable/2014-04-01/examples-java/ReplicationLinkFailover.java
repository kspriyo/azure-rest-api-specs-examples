import com.azure.core.util.Context;

/** Samples for ReplicationLinks Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ReplicationLinkFailover.json
     */
    /**
     * Sample code: Failover a replication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void failoverAReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getReplicationLinks()
            .failover(
                "sqlcrudtest-8931", "sqlcrudtest-2137", "testdb", "f0550bf5-07ce-4270-8e4b-71737975973a", Context.NONE);
    }
}
