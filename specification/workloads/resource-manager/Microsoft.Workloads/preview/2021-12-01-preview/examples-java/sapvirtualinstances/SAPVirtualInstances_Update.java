import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.workloads.models.SapVirtualInstance;
import com.azure.resourcemanager.workloads.models.UserAssignedServiceIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Update.json
     */
    /**
     * Sample code: SAPVirtualInstances_Update.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesUpdate(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        SapVirtualInstance resource =
            manager.sapVirtualInstances().getByResourceGroupWithResponse("test-rg", "X00", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("key1", "svi1"))
            .withIdentity(new UserAssignedServiceIdentity().withType(ManagedServiceIdentityType.NONE))
            .apply();
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
