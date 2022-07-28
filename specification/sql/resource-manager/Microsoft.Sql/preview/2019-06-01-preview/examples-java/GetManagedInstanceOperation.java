import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for ManagedInstanceOperations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/GetManagedInstanceOperation.json
     */
    /**
     * Sample code: Gets the managed instance management operation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheManagedInstanceManagementOperation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceOperations()
            .getWithResponse(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                UUID.fromString("00000000-1111-2222-3333-444444444444"),
                Context.NONE);
    }
}
