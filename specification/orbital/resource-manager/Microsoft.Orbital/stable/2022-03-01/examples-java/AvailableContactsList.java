import com.azure.core.util.Context;
import com.azure.resourcemanager.orbital.models.ContactParameters;
import com.azure.resourcemanager.orbital.models.ContactParametersContactProfile;
import java.time.OffsetDateTime;

/** Samples for Spacecrafts ListAvailableContacts. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableContactsList.json
     */
    /**
     * Sample code: List of Contact.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfContact(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager
            .spacecrafts()
            .listAvailableContacts(
                "contoso-Rgp",
                "CONTOSO_SAT",
                new ContactParameters()
                    .withContactProfile(
                        new ContactParametersContactProfile()
                            .withId(
                                "/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"))
                    .withGroundStationName("EASTUS2_0")
                    .withStartTime(OffsetDateTime.parse("2022-03-01T11:30:00Z"))
                    .withEndTime(OffsetDateTime.parse("2022-03-02T11:30:00Z")),
                Context.NONE);
    }
}
