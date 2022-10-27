import com.azure.core.util.Context;

/** Samples for InboundNatRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/InboundNatRuleList.json
     */
    /**
     * Sample code: InboundNatRuleList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void inboundNatRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getInboundNatRules().list("testrg", "lb1", Context.NONE);
    }
}
