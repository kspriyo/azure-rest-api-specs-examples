import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;

/** Samples for TransparentDataEncryptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseTransparentDataEncryptionGet.json
     */
    /**
     * Sample code: Get a database's transparent data encryption configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSTransparentDataEncryptionConfiguration(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getTransparentDataEncryptions()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                TransparentDataEncryptionName.CURRENT,
                Context.NONE);
    }
}
