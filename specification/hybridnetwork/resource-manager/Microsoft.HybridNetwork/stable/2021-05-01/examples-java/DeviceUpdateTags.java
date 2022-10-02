import com.azure.core.util.Context;
import com.azure.resourcemanager.hybridnetwork.models.Device;
import java.util.HashMap;
import java.util.Map;

/** Samples for Devices UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceUpdateTags.json
     */
    /**
     * Sample code: Update hybrid network device tags.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void updateHybridNetworkDeviceTags(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        Device resource =
            manager.devices().getByResourceGroupWithResponse("rg1", "TestDevice", Context.NONE).getValue();
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
