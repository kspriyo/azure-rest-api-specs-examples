import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.InstanceFailoverGroupInner;
import com.azure.resourcemanager.sql.models.InstanceFailoverGroupReadOnlyEndpoint;
import com.azure.resourcemanager.sql.models.InstanceFailoverGroupReadWriteEndpoint;
import com.azure.resourcemanager.sql.models.ManagedInstancePairInfo;
import com.azure.resourcemanager.sql.models.PartnerRegionInfo;
import com.azure.resourcemanager.sql.models.ReadOnlyEndpointFailoverPolicy;
import com.azure.resourcemanager.sql.models.ReadWriteEndpointFailoverPolicy;
import java.util.Arrays;

/** Samples for InstanceFailoverGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/InstanceFailoverGroupCreateOrUpdate.json
     */
    /**
     * Sample code: Create failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getInstanceFailoverGroups()
            .createOrUpdate(
                "Default",
                "Japan East",
                "failover-group-test-3",
                new InstanceFailoverGroupInner()
                    .withReadWriteEndpoint(
                        new InstanceFailoverGroupReadWriteEndpoint()
                            .withFailoverPolicy(ReadWriteEndpointFailoverPolicy.AUTOMATIC)
                            .withFailoverWithDataLossGracePeriodMinutes(480))
                    .withReadOnlyEndpoint(
                        new InstanceFailoverGroupReadOnlyEndpoint()
                            .withFailoverPolicy(ReadOnlyEndpointFailoverPolicy.DISABLED))
                    .withPartnerRegions(Arrays.asList(new PartnerRegionInfo().withLocation("Japan West")))
                    .withManagedInstancePairs(
                        Arrays
                            .asList(
                                new ManagedInstancePairInfo()
                                    .withPrimaryManagedInstanceId(
                                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance")
                                    .withPartnerManagedInstanceId(
                                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance"))),
                Context.NONE);
    }
}
