import com.azure.core.util.Context;

/** Samples for CollectionPartition ListMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBCollectionPartitionGetMetrics.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountRegionGetMetrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountRegionGetMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCollectionPartitions()
            .listMetrics(
                "rg1",
                "ddb1",
                "databaseRid",
                "collectionRid",
                "$filter=(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq"
                    + " '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z",
                Context.NONE);
    }
}
