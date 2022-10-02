import com.azure.core.util.Context;
import com.azure.resourcemanager.hybridnetwork.models.NetworkFunction;
import java.util.HashMap;
import java.util.Map;

/** Samples for NetworkFunctions UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionUpdateTags.json
     */
    /**
     * Sample code: Update tags for network function resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void updateTagsForNetworkFunctionResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        NetworkFunction resource =
            manager.networkFunctions().getByResourceGroupWithResponse("rg", "testNf", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
