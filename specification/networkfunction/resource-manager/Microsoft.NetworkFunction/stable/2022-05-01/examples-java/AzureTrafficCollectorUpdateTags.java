import com.azure.core.util.Context;
import com.azure.resourcemanager.networkfunction.models.AzureTrafficCollector;
import java.util.HashMap;
import java.util.Map;

/** Samples for AzureTrafficCollectors UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorUpdateTags.json
     */
    /**
     * Sample code: Update Traffic Collector tags.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void updateTrafficCollectorTags(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        AzureTrafficCollector resource =
            manager.azureTrafficCollectors().getByResourceGroupWithResponse("rg1", "atc", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "value1", "key2", "value2")).apply();
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
