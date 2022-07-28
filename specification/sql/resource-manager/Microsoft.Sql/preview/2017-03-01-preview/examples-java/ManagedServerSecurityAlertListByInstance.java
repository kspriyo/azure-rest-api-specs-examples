import com.azure.core.util.Context;

/** Samples for ManagedServerSecurityAlertPolicies ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ManagedServerSecurityAlertListByInstance.json
     */
    /**
     * Sample code: Get the managed server's threat detection policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheManagedServerSThreatDetectionPolicies(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedServerSecurityAlertPolicies()
            .listByInstance("securityalert-4799", "securityalert-6440", Context.NONE);
    }
}
