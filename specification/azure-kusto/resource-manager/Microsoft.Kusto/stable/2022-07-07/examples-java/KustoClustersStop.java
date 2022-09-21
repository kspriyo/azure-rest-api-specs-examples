import com.azure.core.util.Context;

/** Samples for Clusters Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClustersStop.json
     */
    /**
     * Sample code: KustoClustersStop.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersStop(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().stop("kustorptest", "kustoCluster2", Context.NONE);
    }
}
