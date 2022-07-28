import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments GetOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get the network endpoints of all outbound dependencies of an App Service Environment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheNetworkEndpointsOfAllOutboundDependenciesOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .getOutboundNetworkDependenciesEndpoints("Sample-WestUSResourceGroup", "SampleAse", Context.NONE);
    }
}
