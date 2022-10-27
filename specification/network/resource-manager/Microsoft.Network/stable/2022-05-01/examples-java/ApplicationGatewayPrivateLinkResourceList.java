import com.azure.core.util.Context;

/** Samples for ApplicationGatewayPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ApplicationGatewayPrivateLinkResourceList.json
     */
    /**
     * Sample code: Lists all private link resources on application gateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllPrivateLinkResourcesOnApplicationGateway(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGatewayPrivateLinkResources()
            .list("rg1", "appgw", Context.NONE);
    }
}
