import com.azure.core.util.Context;

/** Samples for SapCentralInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPCentralInstances_List.json
     */
    /**
     * Sample code: SAPCentralInstances_List.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPCentralInstancesList(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapCentralInstances().list("test-rg", "X00", Context.NONE);
    }
}
