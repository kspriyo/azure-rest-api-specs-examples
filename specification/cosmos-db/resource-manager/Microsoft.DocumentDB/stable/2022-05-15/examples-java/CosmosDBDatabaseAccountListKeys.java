import com.azure.core.util.Context;

/** Samples for DatabaseAccounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDatabaseAccountListKeys.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountListKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .listKeysWithResponse("rg1", "ddb1", Context.NONE);
    }
}
