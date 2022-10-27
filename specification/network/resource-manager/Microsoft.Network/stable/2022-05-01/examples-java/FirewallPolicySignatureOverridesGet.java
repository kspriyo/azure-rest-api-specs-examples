import com.azure.core.util.Context;

/** Samples for FirewallPolicyIdpsSignaturesOverrides Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicySignatureOverridesGet.json
     */
    /**
     * Sample code: get signature overrides.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSignatureOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyIdpsSignaturesOverrides()
            .getWithResponse("rg1", "firewallPolicy", Context.NONE);
    }
}
