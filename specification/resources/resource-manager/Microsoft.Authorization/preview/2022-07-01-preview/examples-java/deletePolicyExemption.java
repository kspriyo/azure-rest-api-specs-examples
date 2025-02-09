
/** Samples for PolicyExemptions Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/
     * deletePolicyExemption.json
     */
    /**
     * Sample code: Delete a policy exemption.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyExemption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyExemptions().deleteWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM",
            com.azure.core.util.Context.NONE);
    }
}
