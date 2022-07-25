const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get specific `Reservation` details.
 *
 * @summary Get specific `Reservation` details.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationDetails.json
 */
async function getReservation() {
  const reservationId = "6ef59113-3482-40da-8d79-787f823e34bc";
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const expand = "renewProperties";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.reservation.get(reservationId, reservationOrderId, options);
  console.log(result);
}

getReservation().catch(console.error);
