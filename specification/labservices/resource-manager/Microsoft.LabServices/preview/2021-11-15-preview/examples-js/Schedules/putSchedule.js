const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to create or update a lab schedule.
 *
 * @summary Operation to create or update a lab schedule.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/putSchedule.json
 */
async function putSchedule() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const scheduleName = "schedule1";
  const body = {
    notes: "Schedule 1 for students",
    recurrencePattern: {
      expirationDate: new Date("2020-08-14"),
      frequency: "Daily",
      interval: 2,
    },
    startAt: new Date("2020-05-26T12:00:00Z"),
    stopAt: new Date("2020-05-26T18:00:00Z"),
    timeZoneId: "America/Los_Angeles",
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.schedules.createOrUpdate(
    resourceGroupName,
    labName,
    scheduleName,
    body
  );
  console.log(result);
}

putSchedule().catch(console.error);
