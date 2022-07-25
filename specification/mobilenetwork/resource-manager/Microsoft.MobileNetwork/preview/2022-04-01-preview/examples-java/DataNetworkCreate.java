/** Samples for DataNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/DataNetworkCreate.json
     */
    /**
     * Sample code: Create data network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createDataNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .dataNetworks()
            .define("testDataNetwork")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withDescription("myFavouriteDataNetwork")
            .create();
    }
}
