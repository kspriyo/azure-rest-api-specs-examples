import com.azure.core.util.Context;

/** Samples for Schedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_Get.json
     */
    /**
     * Sample code: Schedules_GetByPool.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesGetByPool(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.schedules().getWithResponse("rg1", "TestProject", "DevPool", "autoShutdown", null, Context.NONE);
    }
}
