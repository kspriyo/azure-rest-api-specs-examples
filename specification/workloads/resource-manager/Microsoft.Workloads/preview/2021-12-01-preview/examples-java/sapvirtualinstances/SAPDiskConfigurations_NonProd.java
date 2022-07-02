import com.azure.core.util.Context;

/** Samples for ResourceProvider SapDiskConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPDiskConfigurations_NonProd.json
     */
    /**
     * Sample code: SAPDiskConfigurations_NonProd.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDiskConfigurationsNonProd(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapDiskConfigurationsWithResponse("centralus", null, Context.NONE);
    }
}
