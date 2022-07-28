import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.TransparentDataEncryptionInner;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionStatus;

/** Samples for TransparentDataEncryptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseTransparentDataEncryptionCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a database's transparent data encryption configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADatabaseSTransparentDataEncryptionConfiguration(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getTransparentDataEncryptions()
            .createOrUpdateWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                TransparentDataEncryptionName.CURRENT,
                new TransparentDataEncryptionInner().withStatus(TransparentDataEncryptionStatus.ENABLED),
                Context.NONE);
    }
}
