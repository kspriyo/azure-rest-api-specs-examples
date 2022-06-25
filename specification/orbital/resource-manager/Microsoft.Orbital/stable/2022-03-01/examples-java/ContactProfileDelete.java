import com.azure.core.util.Context;

/** Samples for ContactProfiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileDelete.json
     */
    /**
     * Sample code: Delete Contact Profile.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void deleteContactProfile(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contactProfiles().delete("contoso-Rgp", "CONTOSO-CP", Context.NONE);
    }
}
