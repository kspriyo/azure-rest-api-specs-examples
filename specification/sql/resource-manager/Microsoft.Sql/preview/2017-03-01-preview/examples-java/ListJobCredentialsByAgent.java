import com.azure.core.util.Context;

/** Samples for JobCredentials ListByAgent. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobCredentialsByAgent.json
     */
    /**
     * Sample code: List credentials in a job agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCredentialsInAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobCredentials()
            .listByAgent("group1", "server1", "agent1", Context.NONE);
    }
}
