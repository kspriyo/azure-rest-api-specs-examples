import com.azure.core.util.Context;

/** Samples for AgentPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-06-01/examples/AgentPoolsList.json
     */
    /**
     * Sample code: List Agent Pools by Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAgentPoolsByManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().list("rg1", "clustername1", Context.NONE);
    }
}
