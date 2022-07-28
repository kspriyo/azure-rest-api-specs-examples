import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.PrivateLinkConnectionApprovalRequestResource;
import com.azure.resourcemanager.appservice.models.PrivateLinkConnectionState;

/** Samples for WebApps ApproveOrRejectPrivateEndpointConnectionSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/ApproveRejectSitePrivateEndpointConnectionSlot.json
     */
    /**
     * Sample code: Approves or rejects a private endpoint connection for a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void approvesOrRejectsAPrivateEndpointConnectionForASite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .approveOrRejectPrivateEndpointConnectionSlot(
                "rg",
                "testSite",
                "connection",
                "stage",
                new PrivateLinkConnectionApprovalRequestResource()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkConnectionState()
                            .withStatus("Approved")
                            .withDescription("Approved by admin.")
                            .withActionsRequired("")),
                Context.NONE);
    }
}
