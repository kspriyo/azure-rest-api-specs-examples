import com.azure.core.util.Context;

/** Samples for Diagnostics ListSiteDetectorResponses. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_ListSiteDetectorResponses.json
     */
    /**
     * Sample code: Get App Detector Responses.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppDetectorResponses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .listSiteDetectorResponses("Sample-WestUSResourceGroup", "SampleApp", Context.NONE);
    }
}
