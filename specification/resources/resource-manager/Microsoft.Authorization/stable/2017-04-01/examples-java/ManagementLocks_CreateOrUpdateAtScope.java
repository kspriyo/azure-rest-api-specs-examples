
import com.azure.resourcemanager.resources.fluent.models.ManagementLockObjectInner;
import com.azure.resourcemanager.resources.models.LockLevel;

/** Samples for ManagementLocks CreateOrUpdateByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_CreateOrUpdateAtScope.json
     */
    /**
     * Sample code: Create management lock at scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createManagementLockAtScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .createOrUpdateByScopeWithResponse("subscriptions/subscriptionId", "testlock",
                new ManagementLockObjectInner().withLevel(LockLevel.READ_ONLY), com.azure.core.util.Context.NONE);
    }
}
