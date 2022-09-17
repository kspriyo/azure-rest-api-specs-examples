import com.azure.core.util.Context;
import com.azure.resourcemanager.dnsresolver.models.ForwardingRule;
import com.azure.resourcemanager.dnsresolver.models.ForwardingRuleState;
import java.util.HashMap;
import java.util.Map;

/** Samples for ForwardingRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Patch.json
     */
    /**
     * Sample code: Update forwarding rule in a DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void updateForwardingRuleInADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        ForwardingRule resource =
            manager
                .forwardingRules()
                .getWithResponse(
                    "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", Context.NONE)
                .getValue();
        resource
            .update()
            .withMetadata(mapOf("additionalProp2", "value2"))
            .withForwardingRuleState(ForwardingRuleState.DISABLED)
            .apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
