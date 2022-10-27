import com.azure.core.util.Context;

/** Samples for LoadBalancerFrontendIpConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerFrontendIPConfigurationGet.json
     */
    /**
     * Sample code: LoadBalancerFrontendIPConfigurationGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerFrontendIPConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancerFrontendIpConfigurations()
            .getWithResponse("testrg", "lb", "frontend", Context.NONE);
    }
}
