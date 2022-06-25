import com.azure.core.util.Context;

/** Samples for OperationsResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/OperationResultsGet.json
     */
    /**
     * Sample code: KustoOperationResultsGet.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void kustoOperationResultsGet(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.operationsResults().get("eastus2", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3", Context.NONE);
    }
}
