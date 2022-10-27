import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NetworkSecurityGroupInner;
import com.azure.resourcemanager.network.fluent.models.SecurityRuleInner;
import com.azure.resourcemanager.network.models.SecurityRuleAccess;
import com.azure.resourcemanager.network.models.SecurityRuleDirection;
import com.azure.resourcemanager.network.models.SecurityRuleProtocol;
import java.util.Arrays;

/** Samples for NetworkSecurityGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkSecurityGroupCreateWithRule.json
     */
    /**
     * Sample code: Create network security group with rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkSecurityGroupWithRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkSecurityGroups()
            .createOrUpdate(
                "rg1",
                "testnsg",
                new NetworkSecurityGroupInner()
                    .withLocation("eastus")
                    .withSecurityRules(
                        Arrays
                            .asList(
                                new SecurityRuleInner()
                                    .withName("rule1")
                                    .withProtocol(SecurityRuleProtocol.ASTERISK)
                                    .withSourcePortRange("*")
                                    .withDestinationPortRange("80")
                                    .withSourceAddressPrefix("*")
                                    .withDestinationAddressPrefix("*")
                                    .withAccess(SecurityRuleAccess.ALLOW)
                                    .withPriority(130)
                                    .withDirection(SecurityRuleDirection.INBOUND))),
                Context.NONE);
    }
}
