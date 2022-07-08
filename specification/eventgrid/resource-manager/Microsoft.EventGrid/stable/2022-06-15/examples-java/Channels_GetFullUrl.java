import com.azure.core.util.Context;

/** Samples for Channels GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Channels_GetFullUrl.json
     */
    /**
     * Sample code: Channels_GetFullUrl.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.channels().getFullUrlWithResponse("examplerg", "examplenamespace", "examplechannel", Context.NONE);
    }
}
