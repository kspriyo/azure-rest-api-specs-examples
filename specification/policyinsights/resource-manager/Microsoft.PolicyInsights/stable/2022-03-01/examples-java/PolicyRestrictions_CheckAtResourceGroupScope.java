import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.policyinsights.models.CheckRestrictionsRequest;
import com.azure.resourcemanager.policyinsights.models.CheckRestrictionsResourceDetails;
import com.azure.resourcemanager.policyinsights.models.PendingField;
import java.io.IOException;
import java.util.Arrays;

/** Samples for PolicyRestrictions CheckAtResourceGroupScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtResourceGroupScope.json
     */
    /**
     * Sample code: Check policy restrictions at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void checkPolicyRestrictionsAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) throws IOException {
        manager
            .policyRestrictions()
            .checkAtResourceGroupScopeWithResponse(
                "vmRg",
                new CheckRestrictionsRequest()
                    .withResourceDetails(
                        new CheckRestrictionsResourceDetails()
                            .withResourceContent(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"type\":\"Microsoft.Compute/virtualMachines\",\"properties\":{\"priority\":\"Spot\"}}",
                                        Object.class,
                                        SerializerEncoding.JSON))
                            .withApiVersion("2019-12-01"))
                    .withPendingFields(
                        Arrays
                            .asList(
                                new PendingField().withField("name").withValues(Arrays.asList("myVMName")),
                                new PendingField()
                                    .withField("location")
                                    .withValues(Arrays.asList("eastus", "westus", "westus2", "westeurope")),
                                new PendingField().withField("tags"))),
                Context.NONE);
    }
}
