import com.azure.core.util.Context;

/** Samples for Attestations DeleteAtResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_DeleteResourceScope.json
     */
    /**
     * Sample code: Delete attestation at individual resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void deleteAttestationAtIndividualResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .attestations()
            .deleteAtResourceWithResponse(
                "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM",
                "790996e6-9871-4b1f-9cd9-ec42cd6ced1e",
                Context.NONE);
    }
}
