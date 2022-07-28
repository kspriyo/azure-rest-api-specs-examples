import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceKeyInner;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/** Samples for ManagedInstanceKeys CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ManagedInstanceKeyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a managed instance key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAManagedInstanceKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceKeys()
            .createOrUpdate(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                "someVault_someKey_01234567890123456789012345678901",
                new ManagedInstanceKeyInner()
                    .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
                    .withUri("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
                Context.NONE);
    }
}
