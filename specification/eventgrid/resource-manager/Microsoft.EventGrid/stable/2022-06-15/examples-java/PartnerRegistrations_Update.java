import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerRegistration;
import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerRegistrations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerRegistrations_Update.json
     */
    /**
     * Sample code: PartnerRegistrations_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        PartnerRegistration resource =
            manager
                .partnerRegistrations()
                .getByResourceGroupWithResponse("examplerg", "examplePartnerRegistrationName1", Context.NONE)
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
