import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.fluent.models.WordpressInstanceResourceInner;
import com.azure.resourcemanager.workloads.models.WordpressVersions;

/** Samples for WordpressInstances CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/WordpressInstances_CreateOrUpdate.json
     */
    /**
     * Sample code: Workloads.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void workloads(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .wordpressInstances()
            .createOrUpdate(
                "test-rg",
                "wp39",
                new WordpressInstanceResourceInner()
                    .withVersion(WordpressVersions.FIVE_FOUR_TWO)
                    .withDatabaseName("wpdb")
                    .withDatabaseUser("wpuser"),
                Context.NONE);
    }
}
