import com.azure.core.util.Context;

/** Samples for TimeSeriesDatabaseConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/TimeSeriesDatabaseConnectionsList_example.json
     */
    /**
     * Sample code: List time series database connections for a DigitalTwins instance.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void listTimeSeriesDatabaseConnectionsForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.timeSeriesDatabaseConnections().list("resRg", "myDigitalTwinsService", Context.NONE);
    }
}
