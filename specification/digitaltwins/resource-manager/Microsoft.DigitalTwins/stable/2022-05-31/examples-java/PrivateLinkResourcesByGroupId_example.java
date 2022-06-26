import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/PrivateLinkResourcesByGroupId_example.json
     */
    /**
     * Sample code: Get the specified private link resource for the given Digital Twin.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getTheSpecifiedPrivateLinkResourceForTheGivenDigitalTwin(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.privateLinkResources().getWithResponse("resRg", "myDigitalTwinsService", "subResource", Context.NONE);
    }
}
