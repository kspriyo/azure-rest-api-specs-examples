import com.azure.core.util.Context;

/** Samples for ApplicationGateways ListAvailableSslOptions. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ApplicationGatewayAvailableSslOptionsGet.json
     */
    /**
     * Sample code: Get Available Ssl Options.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableSslOptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .listAvailableSslOptionsWithResponse(Context.NONE);
    }
}
