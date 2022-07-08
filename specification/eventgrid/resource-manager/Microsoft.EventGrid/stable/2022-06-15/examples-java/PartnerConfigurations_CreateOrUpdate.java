import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.fluent.models.PartnerConfigurationInner;
import com.azure.resourcemanager.eventgrid.models.Partner;
import com.azure.resourcemanager.eventgrid.models.PartnerAuthorization;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.UUID;

/** Samples for PartnerConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerConfigurations_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerConfigurations()
            .createOrUpdate(
                "examplerg",
                new PartnerConfigurationInner()
                    .withPartnerAuthorization(
                        new PartnerAuthorization()
                            .withDefaultMaximumExpirationTimeInDays(10)
                            .withAuthorizedPartnersList(
                                Arrays
                                    .asList(
                                        new Partner()
                                            .withPartnerRegistrationImmutableId(
                                                UUID.fromString("941892bc-f5d0-4d1c-8fb5-477570fc2b71"))
                                            .withPartnerName("Contoso.Finance")
                                            .withAuthorizationExpirationTimeInUtc(
                                                OffsetDateTime.parse("2022-01-28T01:20:55.142Z")),
                                        new Partner()
                                            .withPartnerRegistrationImmutableId(
                                                UUID.fromString("5362bdb6-ce3e-4d0d-9a5b-3eb92c8aab38"))
                                            .withPartnerName("fabrikam.HR")
                                            .withAuthorizationExpirationTimeInUtc(
                                                OffsetDateTime.parse("2022-02-20T01:00:00.142Z"))))),
                Context.NONE);
    }
}
