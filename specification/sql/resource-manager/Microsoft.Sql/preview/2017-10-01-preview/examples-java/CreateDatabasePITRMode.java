import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;
import java.time.OffsetDateTime;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/CreateDatabasePITRMode.json
     */
    /**
     * Sample code: Creates a database from PointInTimeRestore.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseFromPointInTimeRestore(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "dbpitr",
                new DatabaseInner()
                    .withLocation("southeastasia")
                    .withSku(new Sku().withName("S0").withTier("Standard"))
                    .withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
                    .withSourceDatabaseId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb")
                    .withRestorePointInTime(OffsetDateTime.parse("2017-07-14T05:35:31.503Z")),
                Context.NONE);
    }
}
