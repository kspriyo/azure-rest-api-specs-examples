import com.azure.core.util.Context;
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.DnsConfig;
import com.azure.resourcemanager.trafficmanager.models.EndpointPropertiesCustomHeadersItem;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfigCustomHeadersItem;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfigExpectedStatusCodeRangesItem;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import com.azure.resourcemanager.trafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.trafficmanager.models.TrafficRoutingMethod;
import com.azure.resourcemanager.trafficmanager.models.TrafficViewEnrollmentStatus;
import java.util.Arrays;

/** Samples for Profiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/Profile-PUT-WithCustomHeaders.json
     */
    /**
     * Sample code: Profile-PUT-WithCustomHeaders.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilePUTWithCustomHeaders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .createOrUpdateWithResponse(
                "azuresdkfornetautoresttrafficmanager2583",
                "azuresdkfornetautoresttrafficmanager6192",
                new ProfileInner()
                    .withLocation("global")
                    .withProfileStatus(ProfileStatus.ENABLED)
                    .withTrafficRoutingMethod(TrafficRoutingMethod.PERFORMANCE)
                    .withDnsConfig(
                        new DnsConfig().withRelativeName("azuresdkfornetautoresttrafficmanager6192").withTtl(35L))
                    .withMonitorConfig(
                        new MonitorConfig()
                            .withProtocol(MonitorProtocol.HTTP)
                            .withPort(80L)
                            .withPath("/testpath.aspx")
                            .withIntervalInSeconds(10L)
                            .withTimeoutInSeconds(5L)
                            .withToleratedNumberOfFailures(2L)
                            .withCustomHeaders(
                                Arrays
                                    .asList(
                                        new MonitorConfigCustomHeadersItem().withName("header-1").withValue("value-1"),
                                        new MonitorConfigCustomHeadersItem().withName("header-2").withValue("value-2")))
                            .withExpectedStatusCodeRanges(
                                Arrays
                                    .asList(
                                        new MonitorConfigExpectedStatusCodeRangesItem().withMin(200).withMax(205),
                                        new MonitorConfigExpectedStatusCodeRangesItem().withMin(400).withMax(410))))
                    .withEndpoints(
                        Arrays
                            .asList(
                                new EndpointInner()
                                    .withName("My external endpoint")
                                    .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints")
                                    .withTarget("foobar.contoso.com")
                                    .withEndpointStatus(EndpointStatus.ENABLED)
                                    .withEndpointLocation("North Europe")
                                    .withCustomHeaders(
                                        Arrays
                                            .asList(
                                                new EndpointPropertiesCustomHeadersItem()
                                                    .withName("header-2")
                                                    .withValue("value-2-overridden")))))
                    .withTrafficViewEnrollmentStatus(TrafficViewEnrollmentStatus.DISABLED),
                Context.NONE);
    }
}
