import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.EventSubscription;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;
import java.util.Arrays;

/** Samples for TopicEventSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/TopicEventSubscriptions_Update.json
     */
    /**
     * Sample code: TopicEventSubscriptions_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        EventSubscription resource =
            manager
                .topicEventSubscriptions()
                .getWithResponse("examplerg", "exampleTopic1", "exampleEventSubscriptionName1", Context.NONE)
                .getValue();
        resource
            .update()
            .withDestination(new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
            .withFilter(
                new EventSubscriptionFilter()
                    .withSubjectBeginsWith("existingPrefix")
                    .withSubjectEndsWith("newSuffix")
                    .withIsSubjectCaseSensitive(true))
            .withLabels(Arrays.asList("label1", "label2"))
            .apply();
    }
}
