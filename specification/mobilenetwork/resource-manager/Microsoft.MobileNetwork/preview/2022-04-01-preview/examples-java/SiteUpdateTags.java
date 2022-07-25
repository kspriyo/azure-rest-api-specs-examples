import com.azure.core.util.Context;
import com.azure.resourcemanager.mobilenetwork.models.Site;
import java.util.HashMap;
import java.util.Map;

/** Samples for Sites UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SiteUpdateTags.json
     */
    /**
     * Sample code: Update mobile network site tags.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void updateMobileNetworkSiteTags(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        Site resource =
            manager.sites().getWithResponse("rg1", "testMobileNetwork", "testSite", Context.NONE).getValue();
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
