
/** Samples for Subscriptions ListLocations. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetLocations.json
     */
    /**
     * Sample code: GetLocationsWithASubscriptionId.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocationsWithASubscriptionId(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().subscriptionClient().getSubscriptions()
            .listLocations("a1ffc958-d2c7-493e-9f1e-125a0477f536", null, com.azure.core.util.Context.NONE);
    }
}
