import com.azure.resourcemanager.mobilenetwork.models.AzureStackEdgeDeviceResourceId;
import com.azure.resourcemanager.mobilenetwork.models.BillingSku;
import com.azure.resourcemanager.mobilenetwork.models.ConnectedClusterResourceId;
import com.azure.resourcemanager.mobilenetwork.models.CoreNetworkType;
import com.azure.resourcemanager.mobilenetwork.models.CustomLocationResourceId;
import com.azure.resourcemanager.mobilenetwork.models.InterfaceProperties;
import com.azure.resourcemanager.mobilenetwork.models.KeyVaultCertificate;
import com.azure.resourcemanager.mobilenetwork.models.LocalDiagnosticsAccessConfiguration;
import com.azure.resourcemanager.mobilenetwork.models.MobileNetworkResourceId;
import com.azure.resourcemanager.mobilenetwork.models.PlatformConfiguration;
import com.azure.resourcemanager.mobilenetwork.models.PlatformType;

/** Samples for PacketCoreControlPlanes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneCreate.json
     */
    /**
     * Sample code: Create packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreControlPlanes()
            .define("TestPacketCoreCP")
            .withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withMobileNetwork(
                new MobileNetworkResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"))
            .withControlPlaneAccessInterface(new InterfaceProperties().withName("N2"))
            .withSku(BillingSku.fromString("testSku"))
            .withPlatform(
                new PlatformConfiguration()
                    .withType(PlatformType.AKS_HCI)
                    .withAzureStackEdgeDevice(
                        new AzureStackEdgeDeviceResourceId()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"))
                    .withConnectedCluster(
                        new ConnectedClusterResourceId()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"))
                    .withCustomLocation(
                        new CustomLocationResourceId()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation")))
            .withCoreNetworkTechnology(CoreNetworkType.FIVE_GC)
            .withVersion("0.2.0")
            .withLocalDiagnosticsAccess(
                new LocalDiagnosticsAccessConfiguration()
                    .withHttpsServerCertificate(
                        new KeyVaultCertificate()
                            .withCertificateUrl("https://contosovault.vault.azure.net/certificates/ingress")))
            .create();
    }
}
