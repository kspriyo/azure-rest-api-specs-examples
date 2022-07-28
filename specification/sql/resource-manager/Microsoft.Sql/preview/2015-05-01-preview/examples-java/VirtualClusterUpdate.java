import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.VirtualClusterUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualClusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/VirtualClusterUpdate.json
     */
    /**
     * Sample code: Update virtual cluster with tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualClusterWithTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getVirtualClusters()
            .update(
                "testrg",
                "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2",
                new VirtualClusterUpdate().withTags(mapOf("tagKey1", "TagValue1")),
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
