import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerAutomaticTuningInner;
import com.azure.resourcemanager.sql.models.AutomaticTuningOptionModeDesired;
import com.azure.resourcemanager.sql.models.AutomaticTuningServerMode;
import com.azure.resourcemanager.sql.models.AutomaticTuningServerOptions;
import java.util.HashMap;
import java.util.Map;

/** Samples for ServerAutomaticTuning Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ServerAutomaticTuningUpdateMax.json
     */
    /**
     * Sample code: Updates server automatic tuning settings with all properties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesServerAutomaticTuningSettingsWithAllProperties(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerAutomaticTunings()
            .updateWithResponse(
                "default-sql-onebox",
                "testsvr11",
                new ServerAutomaticTuningInner()
                    .withDesiredState(AutomaticTuningServerMode.AUTO)
                    .withOptions(
                        mapOf(
                            "createIndex",
                            new AutomaticTuningServerOptions().withDesiredState(AutomaticTuningOptionModeDesired.OFF),
                            "dropIndex",
                            new AutomaticTuningServerOptions().withDesiredState(AutomaticTuningOptionModeDesired.ON),
                            "forceLastGoodPlan",
                            new AutomaticTuningServerOptions()
                                .withDesiredState(AutomaticTuningOptionModeDesired.DEFAULT))),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
