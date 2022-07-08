import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerTopic;
import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerTopics Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopics_Update.json
     */
    /**
     * Sample code: PartnerTopics_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        PartnerTopic resource =
            manager
                .partnerTopics()
                .getByResourceGroupWithResponse("examplerg", "examplePartnerTopicName1", Context.NONE)
                .getValue();
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
