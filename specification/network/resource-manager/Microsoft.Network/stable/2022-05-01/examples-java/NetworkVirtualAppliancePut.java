import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NetworkVirtualApplianceInner;
import com.azure.resourcemanager.network.models.ManagedServiceIdentity;
import com.azure.resourcemanager.network.models.ManagedServiceIdentityUserAssignedIdentities;
import com.azure.resourcemanager.network.models.ResourceIdentityType;
import com.azure.resourcemanager.network.models.VirtualApplianceSkuProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for NetworkVirtualAppliances CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkVirtualAppliancePut.json
     */
    /**
     * Sample code: Create NetworkVirtualAppliance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkVirtualAppliance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkVirtualAppliances()
            .createOrUpdate(
                "rg1",
                "nva",
                new NetworkVirtualApplianceInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withIdentity(
                        new ManagedServiceIdentity()
                            .withType(ResourceIdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                                    new ManagedServiceIdentityUserAssignedIdentities())))
                    .withNvaSku(
                        new VirtualApplianceSkuProperties()
                            .withVendor("Cisco SDWAN")
                            .withBundledScaleUnit("1")
                            .withMarketPlaceVersion("12.1"))
                    .withBootStrapConfigurationBlobs(
                        Arrays
                            .asList(
                                "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"))
                    .withVirtualHub(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"))
                    .withCloudInitConfigurationBlobs(
                        Arrays
                            .asList(
                                "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"))
                    .withVirtualApplianceAsn(10000L),
                Context.NONE);
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
