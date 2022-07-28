import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.JobAgentUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for JobAgents Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/UpdateJobAgent.json
     */
    /**
     * Sample code: Update a job agent's tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAJobAgentSTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobAgents()
            .update(
                "group1",
                "server1",
                "agent1",
                new JobAgentUpdate().withTags(mapOf("mytag1", "myvalue1")),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
