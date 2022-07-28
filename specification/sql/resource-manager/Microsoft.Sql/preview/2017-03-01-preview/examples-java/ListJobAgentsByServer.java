import com.azure.core.util.Context;

/** Samples for JobAgents ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobAgentsByServer.json
     */
    /**
     * Sample code: List job agents in a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobAgentsInAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobAgents().listByServer("group1", "server1", Context.NONE);
    }
}
