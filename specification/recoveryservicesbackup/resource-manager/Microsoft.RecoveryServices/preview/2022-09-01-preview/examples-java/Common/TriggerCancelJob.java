import com.azure.core.util.Context;

/** Samples for JobCancellations Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/TriggerCancelJob.json
     */
    /**
     * Sample code: Cancel Job.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void cancelJob(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .jobCancellations()
            .triggerWithResponse(
                "NetSDKTestRsVault", "SwaggerTestRg", "00000000-0000-0000-0000-000000000000", Context.NONE);
    }
}
