import com.azure.resourcemanager.orbital.models.Direction;
import com.azure.resourcemanager.orbital.models.Polarization;
import com.azure.resourcemanager.orbital.models.SpacecraftLink;
import java.util.Arrays;

/** Samples for Spacecrafts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftCreate.json
     */
    /**
     * Sample code: Create a spacecraft.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void createASpacecraft(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager
            .spacecrafts()
            .define("CONTOSO_SAT")
            .withRegion("eastus2")
            .withExistingResourceGroup("contoso-Rgp")
            .withNoradId("36411")
            .withTitleLine("CONTOSO_SAT")
            .withTleLine1("1 27424U 02022A   22167.05119303  .00000638  00000+0  15103-3 0  9994")
            .withTleLine2("2 27424  98.2477 108.9546 0000928  92.9194 327.0802 14.57300770 69982")
            .withLinks(
                Arrays
                    .asList(
                        new SpacecraftLink()
                            .withName("uplink_lhcp1")
                            .withCenterFrequencyMHz(2250f)
                            .withBandwidthMHz(2f)
                            .withDirection(Direction.UPLINK)
                            .withPolarization(Polarization.LHCP),
                        new SpacecraftLink()
                            .withName("downlink_rhcp1")
                            .withCenterFrequencyMHz(8160f)
                            .withBandwidthMHz(15f)
                            .withDirection(Direction.DOWNLINK)
                            .withPolarization(Polarization.RHCP)))
            .create();
    }
}
