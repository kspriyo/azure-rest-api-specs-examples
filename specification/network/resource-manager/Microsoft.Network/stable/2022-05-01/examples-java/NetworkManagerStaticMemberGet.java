import com.azure.core.util.Context;

/** Samples for StaticMembers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerStaticMemberGet.json
     */
    /**
     * Sample code: StaticMembersGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticMembersGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getStaticMembers()
            .getWithResponse("rg1", "testNetworkManager", "testNetworkGroup", "testStaticMember", Context.NONE);
    }
}
