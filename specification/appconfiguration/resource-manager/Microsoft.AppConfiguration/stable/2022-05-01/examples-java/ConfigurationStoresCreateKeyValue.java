import java.util.HashMap;
import java.util.Map;

/** Samples for KeyValues CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresCreateKeyValue.json
     */
    /**
     * Sample code: KeyValues_CreateOrUpdate.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesCreateOrUpdate(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .keyValues()
            .define("myKey$myLabel")
            .withExistingConfigurationStore("myResourceGroup", "contoso")
            .withTags(mapOf("tag1", "tagValue1", "tag2", "tagValue2"))
            .withValue("myValue")
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
