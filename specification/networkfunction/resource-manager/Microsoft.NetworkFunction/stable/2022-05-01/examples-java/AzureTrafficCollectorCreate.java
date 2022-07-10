import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AzureTrafficCollectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorCreate.json
     */
    /**
     * Sample code: Create a traffic collector.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void createATrafficCollector(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager
            .azureTrafficCollectors()
            .define("atc")
            .withRegion("West US")
            .withExistingResourceGroup("rg1")
            .withTags(mapOf("key1", "value1"))
            .withCollectorPolicies(Arrays.asList())
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
