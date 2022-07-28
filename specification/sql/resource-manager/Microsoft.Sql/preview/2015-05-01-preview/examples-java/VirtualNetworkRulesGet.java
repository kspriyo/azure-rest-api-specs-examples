import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/VirtualNetworkRulesGet.json
     */
    /**
     * Sample code: Gets a virtual network rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAVirtualNetworkRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getVirtualNetworkRules()
            .getWithResponse("Default", "vnet-test-svr", "vnet-firewall-rule", Context.NONE);
    }
}
