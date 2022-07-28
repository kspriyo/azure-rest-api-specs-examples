import com.azure.core.util.Context;
import java.time.OffsetDateTime;

/** Samples for JobExecutions ListByAgent. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobExecutionsByAgentWithFilter.json
     */
    /**
     * Sample code: List all job executions in a job agent with filtering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllJobExecutionsInAJobAgentWithFiltering(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobExecutions()
            .listByAgent(
                "group1",
                "server1",
                "agent1",
                OffsetDateTime.parse("2017-03-21T19:00:00.0000000Z"),
                OffsetDateTime.parse("2017-03-21T19:05:00.0000000Z"),
                OffsetDateTime.parse("2017-03-21T19:20:00.0000000Z"),
                OffsetDateTime.parse("2017-03-21T19:25:00.0000000Z"),
                false,
                null,
                null,
                Context.NONE);
    }
}
