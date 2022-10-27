import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NatGatewayInner;
import com.azure.resourcemanager.network.models.NatGatewaySku;
import com.azure.resourcemanager.network.models.NatGatewaySkuName;
import java.util.Arrays;

/** Samples for NatGateways CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NatGatewayCreateOrUpdate.json
     */
    /**
     * Sample code: Create nat gateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNatGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNatGateways()
            .createOrUpdate(
                "rg1",
                "test-natgateway",
                new NatGatewayInner()
                    .withLocation("westus")
                    .withSku(new NatGatewaySku().withName(NatGatewaySkuName.STANDARD))
                    .withPublicIpAddresses(
                        Arrays
                            .asList(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1")))
                    .withPublicIpPrefixes(
                        Arrays
                            .asList(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"))),
                Context.NONE);
    }
}
