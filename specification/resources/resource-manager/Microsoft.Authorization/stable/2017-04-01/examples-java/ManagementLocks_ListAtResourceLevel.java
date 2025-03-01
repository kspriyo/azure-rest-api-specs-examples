
/** Samples for ManagementLocks ListAtResourceLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_ListAtResourceLevel.json
     */
    /**
     * Sample code: List management locks at resource level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagementLocksAtResourceLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks().listAtResourceLevel(
            "resourcegroupname", "Microsoft.Storage", "parentResourcePath", "storageAccounts", "teststorageaccount",
            null, com.azure.core.util.Context.NONE);
    }
}
