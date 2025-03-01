
/** Samples for ManagementLocks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_ListAtSubscriptionLevel.json
     */
    /**
     * Sample code: List management locks at subscription level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagementLocksAtSubscriptionLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks().list(null,
            com.azure.core.util.Context.NONE);
    }
}
