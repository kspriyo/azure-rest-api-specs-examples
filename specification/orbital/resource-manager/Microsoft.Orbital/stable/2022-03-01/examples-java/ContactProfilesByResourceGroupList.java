import com.azure.core.util.Context;

/** Samples for ContactProfiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfilesByResourceGroupList.json
     */
    /**
     * Sample code: List of Contact Profiles by Resource Group.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfContactProfilesByResourceGroup(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contactProfiles().listByResourceGroup("contoso-Rgp", "opaqueString", Context.NONE);
    }
}
