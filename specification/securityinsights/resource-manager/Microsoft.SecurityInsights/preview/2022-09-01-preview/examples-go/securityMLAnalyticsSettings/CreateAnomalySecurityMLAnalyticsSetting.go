package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/securityMLAnalyticsSettings/CreateAnomalySecurityMLAnalyticsSetting.json
func ExampleSecurityMLAnalyticsSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewSecurityMLAnalyticsSettingsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "f209187f-1d17-4431-94af-c141bf5f23db", &armsecurityinsights.AnomalySecurityMLAnalyticsSettings{
		Etag: to.Ptr("\"260090e2-0000-0d00-0000-5d6fb8670000\""),
		Kind: to.Ptr(armsecurityinsights.SecurityMLAnalyticsSettingsKindAnomaly),
		Properties: &armsecurityinsights.AnomalySecurityMLAnalyticsSettingsProperties{
			Description:            to.Ptr("When account logs from a source region that has rarely been logged in from during the last 14 days, an anomaly is triggered."),
			AnomalySettingsVersion: to.Ptr[int32](0),
			AnomalyVersion:         to.Ptr("1.0.5"),
			CustomizableObservations: map[string]interface{}{
				"multiSelectObservations":       nil,
				"prioritizeExcludeObservations": nil,
				"singleSelectObservations": []interface{}{
					map[string]interface{}{
						"name":           "Device vendor",
						"description":    "Select device vendor of network connection logs from CommonSecurityLog",
						"rerun":          "RerunAlways",
						"sequenceNumber": float64(1),
						"supportedValues": []interface{}{
							"Palo Alto Networks",
							"Fortinet",
							"Check Point",
						},
						"supportedValuesKql": nil,
						"value": []interface{}{
							"Palo Alto Networks",
						},
						"valuesKql": nil,
					},
				},
				"singleValueObservations": nil,
				"thresholdObservations": []interface{}{
					map[string]interface{}{
						"name":           "Daily data transfer threshold in MB",
						"description":    "Suppress anomalies when daily data transfered (in MB) per hour is less than the chosen value",
						"maximum":        "100",
						"minimum":        "1",
						"rerun":          "RerunAlways",
						"sequenceNumber": float64(1),
						"value":          "25",
					},
					map[string]interface{}{
						"name":           "Number of standard deviations",
						"description":    "Triggers anomalies when number of standard deviations is greater than the chosen value",
						"maximum":        "10",
						"minimum":        "2",
						"rerun":          "RerunAlways",
						"sequenceNumber": float64(2),
						"value":          "3",
					},
				},
			},
			DisplayName:       to.Ptr("Login from unusual region"),
			Enabled:           to.Ptr(true),
			Frequency:         to.Ptr("PT1H"),
			IsDefaultSettings: to.Ptr(true),
			RequiredDataConnectors: []*armsecurityinsights.SecurityMLAnalyticsSettingsDataSource{
				{
					ConnectorID: to.Ptr("AWS"),
					DataTypes: []*string{
						to.Ptr("AWSCloudTrail")},
				}},
			SettingsDefinitionID: to.Ptr("f209187f-1d17-4431-94af-c141bf5f23db"),
			SettingsStatus:       to.Ptr(armsecurityinsights.SettingsStatusProduction),
			Tactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticExfiltration),
				to.Ptr(armsecurityinsights.AttackTacticCommandAndControl)},
			Techniques: []*string{
				to.Ptr("T1037"),
				to.Ptr("T1021")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
