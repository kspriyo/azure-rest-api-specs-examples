import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.AppServiceCertificatePatchResourceInner;

/** Samples for AppServiceCertificateOrders UpdateCertificate. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/UpdateAppServiceCertificate.json
     */
    /**
     * Sample code: Update Certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceCertificateOrders()
            .updateCertificateWithResponse(
                "testrg123",
                "SampleCertificateOrderName",
                "SampleCertName1",
                new AppServiceCertificatePatchResourceInner()
                    .withKeyVaultId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName")
                    .withKeyVaultSecretName("SampleSecretName1"),
                Context.NONE);
    }
}
