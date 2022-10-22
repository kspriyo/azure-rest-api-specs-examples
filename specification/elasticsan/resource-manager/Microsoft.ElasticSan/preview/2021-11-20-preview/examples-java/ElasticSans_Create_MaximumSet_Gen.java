import com.azure.resourcemanager.elasticsan.models.Sku;
import com.azure.resourcemanager.elasticsan.models.SkuName;
import com.azure.resourcemanager.elasticsan.models.SkuTier;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ElasticSans Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Create_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansCreateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .elasticSans()
            .define("ti7q-k952-1qB3J_5")
            .withRegion("aaaaaaaaaaaaaaaaaaaaaaaaaaa")
            .withExistingResourceGroup("rgelasticsan")
            .withSku(new Sku().withName(SkuName.PREMIUM_LRS).withTier(SkuTier.PREMIUM))
            .withBaseSizeTiB(26L)
            .withExtendedCapacitySizeTiB(7L)
            .withTags(mapOf("key896", "aaaaaaaaaaaaaaaaaa"))
            .withAvailabilityZones(Arrays.asList("aaaaaaaaaaaaaaaaa"))
            .create();
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
