import com.azure.core.util.Context;

/** Samples for ReplicationLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ReplicationLinkDelete.json
     */
    /**
     * Sample code: Delete a replication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getReplicationLinks()
            .deleteWithResponse(
                "sqlcrudtest-4799", "sqlcrudtest-6440", "testdb", "5b301b68-03f6-4b26-b0f4-73ebb8634238", Context.NONE);
    }
}
