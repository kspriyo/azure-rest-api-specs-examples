const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Calculates price for exchanging `Reservations` if there are no policy errors.

 *
 * @summary Calculates price for exchanging `Reservations` if there are no policy errors.

 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/CalculateExchange.json
 */
async function calculateExchange() {
  const body = {
    properties: {
      reservationsToExchange: [
        {
          quantity: 1,
          reservationId:
            "/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6",
        },
      ],
      reservationsToPurchase: [
        {
          appliedScopeType: "Shared",
          appliedScopes: [],
          billingPlan: "Upfront",
          billingScopeId: "/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83",
          displayName: "testDisplayName",
          location: "westus",
          quantity: 1,
          renew: false,
          reservedResourceProperties: { instanceFlexibility: "On" },
          reservedResourceType: "VirtualMachines",
          sku: { name: "Standard_B1ls" },
          term: "P1Y",
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.calculateExchange.beginPostAndWait(body);
  console.log(result);
}

calculateExchange().catch(console.error);
