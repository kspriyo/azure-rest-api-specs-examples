import com.azure.resourcemanager.devcenter.models.GitCatalog;

/** Samples for Catalogs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Catalogs_CreateGitHub.json
     */
    /**
     * Sample code: Catalogs_CreateOrUpdateGitHub.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsCreateOrUpdateGitHub(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .catalogs()
            .define("{catalogName}")
            .withExistingDevcenter("rg1", "Contoso")
            .withGitHub(
                new GitCatalog()
                    .withUri("https://github.com/Contoso/centralrepo-fake.git")
                    .withBranch("main")
                    .withSecretIdentifier("https://contosokv.vault.azure.net/secrets/CentralRepoPat")
                    .withPath("/templates"))
            .create();
    }
}
