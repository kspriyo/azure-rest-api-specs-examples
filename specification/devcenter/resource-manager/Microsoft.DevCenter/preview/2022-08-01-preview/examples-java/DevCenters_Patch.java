import com.azure.core.util.Context;
import com.azure.resourcemanager.devcenter.models.DevCenter;
import java.util.HashMap;
import java.util.Map;

/** Samples for DevCenters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevCenters_Patch.json
     */
    /**
     * Sample code: DevCenters_Update.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        DevCenter resource =
            manager.devCenters().getByResourceGroupWithResponse("rg1", "Contoso", Context.NONE).getValue();
        resource.update().withTags(mapOf("CostCode", "12345")).apply();
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
