import com.azure.core.util.Context;

/** Samples for NetAppResource QueryRegionInfo. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/RegionInfo.json
     */
    /**
     * Sample code: RegionInfo_Query.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void regionInfoQuery(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResources().queryRegionInfoWithResponse("eastus", Context.NONE);
    }
}
