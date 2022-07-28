import com.azure.core.util.Context;

/** Samples for Service Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDataTransferServiceGet.json
     */
    /**
     * Sample code: DataTransferServiceGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dataTransferServiceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .getWithResponse("rg1", "ddb1", "DataTransfer", Context.NONE);
    }
}
