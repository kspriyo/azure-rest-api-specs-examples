import com.azure.core.util.Context;

/** Samples for NatRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NatRuleGet.json
     */
    /**
     * Sample code: NatRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void natRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNatRules()
            .getWithResponse("rg1", "gateway1", "natRule1", Context.NONE);
    }
}
