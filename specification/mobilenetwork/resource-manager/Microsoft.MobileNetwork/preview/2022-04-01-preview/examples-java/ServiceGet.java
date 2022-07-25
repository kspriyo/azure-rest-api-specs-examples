import com.azure.core.util.Context;

/** Samples for Services Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/ServiceGet.json
     */
    /**
     * Sample code: Get service.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getService(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.services().getWithResponse("rg1", "testMobileNetwork", "TestService", Context.NONE);
    }
}
