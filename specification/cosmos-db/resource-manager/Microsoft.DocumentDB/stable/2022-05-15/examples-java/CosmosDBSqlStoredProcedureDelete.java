import com.azure.core.util.Context;

/** Samples for SqlResources DeleteSqlStoredProcedure. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlStoredProcedureDelete.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlStoredProcedureDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .deleteSqlStoredProcedure(
                "rg1", "ddb1", "databaseName", "containerName", "storedProcedureName", Context.NONE);
    }
}
