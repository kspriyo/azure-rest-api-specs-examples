import com.azure.resourcemanager.mobilenetwork.models.AttachedDataNetworkResourceId;
import com.azure.resourcemanager.mobilenetwork.models.SimPolicyResourceId;
import com.azure.resourcemanager.mobilenetwork.models.SimStaticIpProperties;
import com.azure.resourcemanager.mobilenetwork.models.SimStaticIpPropertiesStaticIp;
import com.azure.resourcemanager.mobilenetwork.models.SliceResourceId;
import java.util.Arrays;

/** Samples for Sims CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimCreate.json
     */
    /**
     * Sample code: Create SIM.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createSIM(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .sims()
            .define("testSim")
            .withExistingSimGroup("rg1", "testSimGroup")
            .withInternationalMobileSubscriberIdentity("00000")
            .withIntegratedCircuitCardIdentifier("8900000000000000000")
            .withAuthenticationKey("00000000000000000000000000000000")
            .withOperatorKeyCode("00000000000000000000000000000000")
            .withDeviceType("Video camera")
            .withSimPolicy(
                new SimPolicyResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"))
            .withStaticIpConfiguration(
                Arrays
                    .asList(
                        new SimStaticIpProperties()
                            .withAttachedDataNetwork(
                                new AttachedDataNetworkResourceId()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"))
                            .withSlice(
                                new SliceResourceId()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"))
                            .withStaticIp(new SimStaticIpPropertiesStaticIp().withIpv4Address("2.4.0.1"))))
            .create();
    }
}
