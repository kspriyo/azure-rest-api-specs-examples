import com.azure.core.util.Context;

/** Samples for AgentPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-06-01/examples/AgentPoolsDelete.json
     */
    /**
     * Sample code: Delete Agent Pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getAgentPools()
            .delete("rg1", "clustername1", "agentpool1", Context.NONE);
    }
}
