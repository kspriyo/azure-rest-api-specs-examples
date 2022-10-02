package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateAPIPolling.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAApiPollingDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8", &armsecurityinsights.CodelessAPIPollingDataConnector{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAPIPolling),
		Properties: &armsecurityinsights.APIPollingParameters{
			ConnectorUIConfig: &armsecurityinsights.CodelessUIConnectorConfigProperties{
				Availability: &armsecurityinsights.Availability{
					IsPreview: to.Ptr(true),
					Status:    to.Ptr[int32](1),
				},
				ConnectivityCriteria: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesConnectivityCriteriaItem{
					{
						Type:  to.Ptr(armsecurityinsights.ConnectivityType("SentinelKindsV2")),
						Value: []*string{},
					}},
				DataTypes: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesDataTypesItem{
					{
						Name:                  to.Ptr("{{graphQueriesTableName}}"),
						LastDataReceivedQuery: to.Ptr("{{graphQueriesTableName}}\n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)"),
					}},
				DescriptionMarkdown: to.Ptr("The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process."),
				GraphQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesGraphQueriesItem{
					{
						BaseQuery:  to.Ptr("{{graphQueriesTableName}}"),
						Legend:     to.Ptr("GitHub audit log events"),
						MetricName: to.Ptr("Total events received"),
					}},
				GraphQueriesTableName: to.Ptr("GitHubAuditLogPolling_CL"),
				InstructionSteps: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesInstructionStepsItem{
					{
						Description: to.Ptr("Enable GitHub audit Logs. \n Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create or find your personal key"),
						Instructions: []*armsecurityinsights.InstructionStepsInstructionsItem{
							{
								Type: to.Ptr(armsecurityinsights.SettingType("APIKey")),
								Parameters: map[string]interface{}{
									"enable": "true",
									"userRequestPlaceHoldersInput": []interface{}{
										map[string]interface{}{
											"displayText":      "Organization Name",
											"placeHolderName":  "{{placeHolder1}}",
											"placeHolderValue": "",
											"requestObjectKey": "apiEndpoint",
										},
									},
								},
							}},
						Title: to.Ptr("Connect GitHub Enterprise Audit Log to Azure Sentinel"),
					}},
				Permissions: &armsecurityinsights.Permissions{
					Customs: []*armsecurityinsights.PermissionsCustomsItem{
						{
							Name:        to.Ptr("GitHub API personal token Key"),
							Description: to.Ptr("You need access to GitHub personal token, the key should have 'admin:org' scope"),
						}},
					ResourceProvider: []*armsecurityinsights.PermissionsResourceProviderItem{
						{
							PermissionsDisplayText: to.Ptr("read and write permissions are required."),
							Provider:               to.Ptr(armsecurityinsights.ProviderNameMicrosoftOperationalInsightsWorkspaces),
							ProviderDisplayName:    to.Ptr("Workspace"),
							RequiredPermissions: &armsecurityinsights.RequiredPermissions{
								Delete: to.Ptr(true),
								Read:   to.Ptr(true),
								Write:  to.Ptr(true),
							},
							Scope: to.Ptr(armsecurityinsights.PermissionProviderScopeWorkspace),
						}},
				},
				Publisher: to.Ptr("GitHub"),
				SampleQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesSampleQueriesItem{
					{
						Description: to.Ptr("All logs"),
						Query:       to.Ptr("{{graphQueriesTableName}}\n | take 10 <change>"),
					}},
				Title: to.Ptr("GitHub Enterprise Audit Log"),
			},
			PollingConfig: &armsecurityinsights.CodelessConnectorPollingConfigProperties{
				Auth: &armsecurityinsights.CodelessConnectorPollingAuthProperties{
					APIKeyIdentifier: to.Ptr("token"),
					APIKeyName:       to.Ptr("Authorization"),
					AuthType:         to.Ptr("APIKey"),
				},
				Paging: &armsecurityinsights.CodelessConnectorPollingPagingProperties{
					PageSizeParaName: to.Ptr("per_page"),
					PagingType:       to.Ptr("LinkHeader"),
				},
				Response: &armsecurityinsights.CodelessConnectorPollingResponseProperties{
					EventsJSONPaths: []*string{
						to.Ptr("$")},
				},
				Request: &armsecurityinsights.CodelessConnectorPollingRequestProperties{
					APIEndpoint: to.Ptr("https://api.github.com/organizations/{{placeHolder1}}/audit-log"),
					Headers: map[string]interface{}{
						"Accept":     "application/json",
						"User-Agent": "Scuba",
					},
					HTTPMethod: to.Ptr("Get"),
					QueryParameters: map[string]interface{}{
						"phrase": "created:{_QueryWindowStartTime}..{_QueryWindowEndTime}",
					},
					QueryTimeFormat:  to.Ptr("yyyy-MM-ddTHH:mm:ssZ"),
					QueryWindowInMin: to.Ptr[int32](15),
					RateLimitQPS:     to.Ptr[int32](50),
					RetryCount:       to.Ptr[int32](2),
					TimeoutInSeconds: to.Ptr[int32](60),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
