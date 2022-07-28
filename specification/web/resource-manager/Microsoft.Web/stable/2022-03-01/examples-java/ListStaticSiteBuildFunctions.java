import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteBuildFunctions. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/ListStaticSiteBuildFunctions.json
     */
    /**
     * Sample code: Gets the functions of a particular static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheFunctionsOfAParticularStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteBuildFunctions("rg", "testStaticSite0", "default", Context.NONE);
    }
}
