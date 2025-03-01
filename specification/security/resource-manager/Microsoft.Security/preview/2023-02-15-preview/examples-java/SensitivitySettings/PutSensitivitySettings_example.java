
import com.azure.resourcemanager.security.models.UpdateSensitivitySettingsRequest;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for ResourceProvider UpdateSensitivitySettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/
     * SensitivitySettings/PutSensitivitySettings_example.json
     */
    /**
     * Sample code: Update sensitivity settings.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updateSensitivitySettings(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.resourceProviders().updateSensitivitySettingsWithResponse(
            new UpdateSensitivitySettingsRequest()
                .withSensitiveInfoTypesIds(Arrays.asList(UUID.fromString("f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb"),
                    UUID.fromString("b452f22b-f87d-4f48-8490-ecf0873325b5"),
                    UUID.fromString("d59ee8b6-2618-404b-a5e7-aa377cd67543")))
                .withSensitivityThresholdLabelOrder(2.0F).withSensitivityThresholdLabelId(
                    UUID.fromString("f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb")),
            com.azure.core.util.Context.NONE);
    }
}
