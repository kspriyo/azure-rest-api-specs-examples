import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuitAuthorizations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCircuitAuthorizationDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit Authorization.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteExpressRouteCircuitAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuitAuthorizations()
            .delete("rg1", "circuitName", "authorizationName", Context.NONE);
    }
}
