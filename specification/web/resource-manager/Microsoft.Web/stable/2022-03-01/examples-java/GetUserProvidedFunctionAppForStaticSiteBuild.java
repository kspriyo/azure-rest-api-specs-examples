import com.azure.core.util.Context;

/** Samples for StaticSites GetUserProvidedFunctionAppForStaticSiteBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetUserProvidedFunctionAppForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getUserProvidedFunctionAppForStaticSiteBuildWithResponse(
                "rg", "testStaticSite0", "default", "testFunctionApp", Context.NONE);
    }
}
