import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.BastionHostInner;
import com.azure.resourcemanager.network.models.BastionHostIpConfiguration;
import java.util.Arrays;

/** Samples for BastionHosts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/BastionHostPut.json
     */
    /**
     * Sample code: Create Bastion Host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createBastionHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getBastionHosts()
            .createOrUpdate(
                "rg1",
                "bastionhosttenant",
                new BastionHostInner()
                    .withIpConfigurations(
                        Arrays
                            .asList(
                                new BastionHostIpConfiguration()
                                    .withName("bastionHostIpConfiguration")
                                    .withSubnet(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"))
                                    .withPublicIpAddress(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName")))),
                Context.NONE);
    }
}
