import com.azure.core.util.Context;
import com.azure.resourcemanager.devhub.models.Workflow;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workflow UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/Workflow_UpdateTags.json
     */
    /**
     * Sample code: Update Managed Cluster Tags.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void updateManagedClusterTags(com.azure.resourcemanager.devhub.DevHubManager manager) {
        Workflow resource =
            manager.workflows().getByResourceGroupWithResponse("resourceGroup1", "workflow1", Context.NONE).getValue();
        resource.update().withTags(mapOf("promote", "false", "resourceEnv", "testing")).apply();
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
