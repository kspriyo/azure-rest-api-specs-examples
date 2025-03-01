
import com.azure.resourcemanager.resources.fluent.models.ManagementLockObjectInner;
import com.azure.resourcemanager.resources.models.LockLevel;

/** Samples for ManagementLocks CreateOrUpdateAtSubscriptionLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_CreateOrUpdateAtSubscriptionLevel.json
     */
    /**
     * Sample code: Create management lock at subscription level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createManagementLockAtSubscriptionLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .createOrUpdateAtSubscriptionLevelWithResponse("testlock",
                new ManagementLockObjectInner().withLevel(LockLevel.READ_ONLY), com.azure.core.util.Context.NONE);
    }
}
