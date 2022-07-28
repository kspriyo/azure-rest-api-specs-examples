import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerCommunicationLinkInner;

/** Samples for ServerCommunicationLinks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkCreateOrUpdate.json
     */
    /**
     * Sample code: Create a server communication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAServerCommunicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerCommunicationLinks()
            .createOrUpdate(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                "link1",
                new ServerCommunicationLinkInner().withPartnerServer("sqldcrudtest-test"),
                Context.NONE);
    }
}
