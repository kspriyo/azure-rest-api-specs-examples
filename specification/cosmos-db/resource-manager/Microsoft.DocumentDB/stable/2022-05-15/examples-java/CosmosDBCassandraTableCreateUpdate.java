import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CassandraPartitionKey;
import com.azure.resourcemanager.cosmos.models.CassandraSchema;
import com.azure.resourcemanager.cosmos.models.CassandraTableCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.CassandraTableResource;
import com.azure.resourcemanager.cosmos.models.ClusterKey;
import com.azure.resourcemanager.cosmos.models.Column;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for CassandraResources CreateUpdateCassandraTable. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBCassandraTableCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBCassandraTableCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .createUpdateCassandraTable(
                "rg1",
                "ddb1",
                "keyspaceName",
                "tableName",
                new CassandraTableCreateUpdateParameters()
                    .withLocation("West US")
                    .withTags(mapOf())
                    .withResource(
                        new CassandraTableResource()
                            .withId("tableName")
                            .withDefaultTtl(100)
                            .withSchema(
                                new CassandraSchema()
                                    .withColumns(Arrays.asList(new Column().withName("columnA").withType("Ascii")))
                                    .withPartitionKeys(Arrays.asList(new CassandraPartitionKey().withName("columnA")))
                                    .withClusterKeys(
                                        Arrays.asList(new ClusterKey().withName("columnA").withOrderBy("Asc")))))
                    .withOptions(new CreateUpdateOptions()),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
