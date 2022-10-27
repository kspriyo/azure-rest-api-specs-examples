import com.azure.core.util.Context;

/** Samples for VipSwap Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CloudServiceSwapGet.json
     */
    /**
     * Sample code: Get swap resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSwapResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVipSwaps()
            .getWithResponse("rg1", "testCloudService", Context.NONE);
    }
}
