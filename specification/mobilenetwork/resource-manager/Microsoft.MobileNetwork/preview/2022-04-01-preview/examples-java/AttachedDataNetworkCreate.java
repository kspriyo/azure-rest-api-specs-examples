import com.azure.resourcemanager.mobilenetwork.models.InterfaceProperties;
import com.azure.resourcemanager.mobilenetwork.models.NaptConfiguration;
import com.azure.resourcemanager.mobilenetwork.models.NaptEnabled;
import com.azure.resourcemanager.mobilenetwork.models.PinholeTimeouts;
import com.azure.resourcemanager.mobilenetwork.models.PortRange;
import com.azure.resourcemanager.mobilenetwork.models.PortReuseHoldTimes;
import java.util.Arrays;

/** Samples for AttachedDataNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/AttachedDataNetworkCreate.json
     */
    /**
     * Sample code: Create attached data network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createAttachedDataNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .attachedDataNetworks()
            .define("TestAttachedDataNetwork")
            .withRegion("eastus")
            .withExistingPacketCoreDataPlane("rg1", "TestPacketCoreCP", "TestPacketCoreDP")
            .withUserPlaneDataInterface(new InterfaceProperties().withName("N6"))
            .withDnsAddresses(Arrays.asList("1.1.1.1"))
            .withNaptConfiguration(
                new NaptConfiguration()
                    .withEnabled(NaptEnabled.ENABLED)
                    .withPortRange(new PortRange().withMinPort(1024).withMaxPort(49999))
                    .withPortReuseHoldTime(new PortReuseHoldTimes().withTcp(120).withUdp(60))
                    .withPinholeLimits(65536)
                    .withPinholeTimeouts(new PinholeTimeouts().withTcp(180).withUdp(30).withIcmp(30)))
            .withUserEquipmentAddressPoolPrefix(Arrays.asList("2.2.0.0/16"))
            .withUserEquipmentStaticAddressPoolPrefix(Arrays.asList("2.4.0.0/16"))
            .create();
    }
}
