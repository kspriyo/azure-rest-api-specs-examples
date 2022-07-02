import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.SapDatabaseInstance;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapDatabaseInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPDatabaseInstances_Update.json
     */
    /**
     * Sample code: SAPDatabaseInstances_Update.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDatabaseInstancesUpdate(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        SapDatabaseInstance resource =
            manager.sapDatabaseInstances().getWithResponse("test-rg", "X00", "databaseServer", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "value1")).apply();
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
