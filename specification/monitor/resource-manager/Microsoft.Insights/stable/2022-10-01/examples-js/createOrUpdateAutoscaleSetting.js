const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an autoscale setting.
 *
 * @summary Creates or updates an autoscale setting.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/createOrUpdateAutoscaleSetting.json
 */
async function createOrUpdateAnAutoscaleSetting() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "TestingMetricsScaleSet";
  const autoscaleSettingName = "MySetting";
  const parameters = {
    enabled: true,
    location: "West US",
    notifications: [
      {
        email: {
          customEmails: ["gu@ms.com", "ge@ns.net"],
          sendToSubscriptionAdministrator: true,
          sendToSubscriptionCoAdministrators: true,
        },
        operation: "Scale",
        webhooks: [{ properties: {}, serviceUri: "http://myservice.com" }],
      },
    ],
    predictiveAutoscalePolicy: { scaleMode: "Enabled" },
    profiles: [
      {
        name: "adios",
        capacity: { default: "1", maximum: "10", minimum: "1" },
        fixedDate: {
          end: new Date("2015-03-05T14:30:00Z"),
          start: new Date("2015-03-05T14:00:00Z"),
          timeZone: "UTC",
        },
        rules: [
          {
            metricTrigger: {
              dividePerInstance: false,
              metricName: "Percentage CPU",
              metricResourceUri:
                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
              operator: "GreaterThan",
              statistic: "Average",
              threshold: 10,
              timeAggregation: "Average",
              timeGrain: "PT1M",
              timeWindow: "PT5M",
            },
            scaleAction: {
              type: "ChangeCount",
              cooldown: "PT5M",
              direction: "Increase",
              value: "1",
            },
          },
          {
            metricTrigger: {
              dividePerInstance: false,
              metricName: "Percentage CPU",
              metricResourceUri:
                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
              operator: "GreaterThan",
              statistic: "Average",
              threshold: 15,
              timeAggregation: "Average",
              timeGrain: "PT2M",
              timeWindow: "PT5M",
            },
            scaleAction: {
              type: "ChangeCount",
              cooldown: "PT6M",
              direction: "Decrease",
              value: "2",
            },
          },
        ],
      },
      {
        name: "saludos",
        capacity: { default: "1", maximum: "10", minimum: "1" },
        recurrence: {
          frequency: "Week",
          schedule: { days: ["1"], hours: [5], minutes: [15], timeZone: "UTC" },
        },
        rules: [
          {
            metricTrigger: {
              dividePerInstance: false,
              metricName: "Percentage CPU",
              metricResourceUri:
                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
              operator: "GreaterThan",
              statistic: "Average",
              threshold: 10,
              timeAggregation: "Average",
              timeGrain: "PT1M",
              timeWindow: "PT5M",
            },
            scaleAction: {
              type: "ChangeCount",
              cooldown: "PT5M",
              direction: "Increase",
              value: "1",
            },
          },
          {
            metricTrigger: {
              dividePerInstance: false,
              metricName: "Percentage CPU",
              metricResourceUri:
                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
              operator: "GreaterThan",
              statistic: "Average",
              threshold: 15,
              timeAggregation: "Average",
              timeGrain: "PT2M",
              timeWindow: "PT5M",
            },
            scaleAction: {
              type: "ChangeCount",
              cooldown: "PT6M",
              direction: "Decrease",
              value: "2",
            },
          },
        ],
      },
    ],
    tags: { key1: "value1", key2: "value2" },
    targetResourceUri:
      "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.autoscaleSettings.createOrUpdate(
    resourceGroupName,
    autoscaleSettingName,
    parameters
  );
  console.log(result);
}

createOrUpdateAnAutoscaleSetting().catch(console.error);
