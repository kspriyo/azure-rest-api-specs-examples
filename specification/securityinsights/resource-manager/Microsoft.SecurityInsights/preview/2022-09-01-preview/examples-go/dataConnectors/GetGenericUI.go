package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetGenericUI.json
func ExampleDataConnectorsClient_Get_getAGenericUiDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.CodelessUIDataConnector{
	// 		Name: to.Ptr("316ec55e-7138-4d63-ab18-90c8a60fd1c8"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/316ec55e-7138-4d63-ab18-90c8a60fd1c8"),
	// 		Etag: to.Ptr("\"1a00b074-0000-0100-0000-606ef5bd0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindGenericUI),
	// 		Properties: &armsecurityinsights.CodelessParameters{
	// 			ConnectorUIConfig: &armsecurityinsights.CodelessUIConnectorConfigProperties{
	// 				Availability: &armsecurityinsights.Availability{
	// 					IsPreview: to.Ptr(true),
	// 					Status: to.Ptr[int32](1),
	// 				},
	// 				ConnectivityCriteria: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesConnectivityCriteriaItem{
	// 					{
	// 						Type: to.Ptr(armsecurityinsights.ConnectivityTypeIsConnectedQuery),
	// 						Value: []*string{
	// 							to.Ptr("{{graphQueriesTableName}}\n            | summarize LastLogReceived = max(TimeGenerated)\n            | project IsConnected = LastLogReceived > ago(30d)")},
	// 					}},
	// 					CustomImage: to.Ptr("The image connector content"),
	// 					DataTypes: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesDataTypesItem{
	// 						{
	// 							Name: to.Ptr("{{graphQueriesTableName}}"),
	// 							LastDataReceivedQuery: to.Ptr("{{graphQueriesTableName}}\n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)"),
	// 					}},
	// 					DescriptionMarkdown: to.Ptr("The [Qualys Vulnerability Management (VM)](https://www.qualys.com/apps/vulnerability-management/) data connector provides the capability to ingest vulnerability host detection data into Azure Sentinel through the Qualys API. The connector provides visibility into host detection data from vulerability scans. This connector provides Azure Sentinel the capability to view dashboards, create custom alerts, and improve investigation "),
	// 					GraphQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesGraphQueriesItem{
	// 						{
	// 							BaseQuery: to.Ptr("{{graphQueriesTableName}}"),
	// 							Legend: to.Ptr("{{graphQueriesTableName}}"),
	// 							MetricName: to.Ptr("Total data received"),
	// 					}},
	// 					GraphQueriesTableName: to.Ptr("QualysHostDetection_CL"),
	// 					InstructionSteps: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesInstructionStepsItem{
	// 						{
	// 							Description: to.Ptr(">**NOTE:** This connector uses Azure Functions to connect to Qualys VM to pull its logs into Azure Sentinel. This might result in additional data ingestion costs. Check the [Azure Functions pricing page](https://azure.microsoft.com/pricing/details/functions/) for details."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr(">**(Optional Step)** Securely store workspace and API authorization key(s) or token(s) in Azure Key Vault. Azure Key Vault provides a secure mechanism to store and retrieve key values. [Follow these instructions](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references) to use Azure Key Vault with an Azure Function App."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("**STEP 1 - Configuration steps for the Qualys VM API**\n\n1. Log into the Qualys Vulnerability Management console with an administrator account, select the **Users** tab and the **Users** subtab. \n2. Click on the **New** drop-down menu and select **Users..**\n3. Create a username and password for the API account. \n4. In the **User Roles** tab, ensure the account role is set to **Manager** and access is allowed to **GUI** and **API**\n4. Log out of the administrator account and log into the console with the new API credentials for validation, then log out of the API account. \n5. Log back into the console using an administrator account and modify the API accounts User Roles, removing access to **GUI**. \n6. Save all changes."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("**STEP 2 - Choose ONE from the following two deployment options to deploy the connector and the associated Azure Function**\n\n>**IMPORTANT:** Before deploying the Qualys VM connector, have the Workspace ID and Workspace Primary Key (can be copied from the following), as well as the Qualys VM API Authorization Key(s), readily available."),
	// 							Instructions: []*armsecurityinsights.InstructionStepsInstructionsItem{
	// 								{
	// 									Type: to.Ptr(armsecurityinsights.SettingTypeCopyableLabel),
	// 									Parameters: map[string]any{
	// 										"fillWith":[]any{
	// 											"WorkspaceId",
	// 										},
	// 										"label": "Workspace ID",
	// 									},
	// 								},
	// 								{
	// 									Type: to.Ptr(armsecurityinsights.SettingTypeCopyableLabel),
	// 									Parameters: map[string]any{
	// 										"fillWith":[]any{
	// 											"PrimaryKey",
	// 										},
	// 										"label": "Primary Key",
	// 									},
	// 							}},
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("Use this method for automated deployment of the Qualys VM connector using an ARM Tempate.\n\n1. Click the **Deploy to Azure** button below. \n\n	[![Deploy To Azure](https://aka.ms/deploytoazurebutton)](https://aka.ms/sentinelqualysvmazuredeploy)\n2. Select the preferred **Subscription**, **Resource Group** and **Location**. \n3. Enter the **Workspace ID**, **Workspace Key**, **API Username**, **API Password** , update the **URI**, and any additional URI **Filter Parameters** (each filter should be separated by an \"&\" symbol, no spaces.) \n> - Enter the URI that corresponds to your region. The complete list of API Server URLs can be [found here](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348) -- There is no need to add a time suffix to the URI, the Function App will dynamically append the Time Value to the URI in the proper format. \n - The default **Time Interval** is set to pull the last five (5) minutes of data. If the time interval needs to be modified, it is recommended to change the Function App Timer Trigger accordingly (in the function.json file, post deployment) to prevent overlapping data ingestion. \n> - Note: If using Azure Key Vault secrets for any of the values above, use the`@Microsoft.KeyVault(SecretUri={Security Identifier})`schema in place of the string values. Refer to [Key Vault references documentation](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references) for further details. \n4. Mark the checkbox labeled **I agree to the terms and conditions stated above**. \n5. Click **Purchase** to deploy."),
	// 							Title: to.Ptr("Option 1 - Azure Resource Manager (ARM) Template"),
	// 						},
	// 						{
	// 							Description: to.Ptr("Use the following step-by-step instructions to deploy the Quayls VM connector manually with Azure Functions."),
	// 							Title: to.Ptr("Option 2 - Manual Deployment of Azure Functions"),
	// 						},
	// 						{
	// 							Description: to.Ptr("**1. Create a Function App**\n\n1.  From the Azure Portal, navigate to [Function App](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Web%2Fsites/kind/functionapp), and select **+ Add**.\n2. In the **Basics** tab, ensure Runtime stack is set to **Powershell Core**. \n3. In the **Hosting** tab, ensure the **Consumption (Serverless)** plan type is selected.\n4. Make other preferrable configuration changes, if needed, then click **Create**."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("**2. Import Function App Code**\n\n1. In the newly created Function App, select **Functions** on the left pane and click **+ New Function**.\n2. Select **Timer Trigger**.\n3. Enter a unique Function **Name** and leave the default cron schedule of every 5 minutes, then click **Create**.\n5. Click on **Code + Test** on the left pane. \n6. Copy the [Function App Code](https://aka.ms/sentinelqualysvmazurefunctioncode) and paste into the Function App `run.ps1` editor.\n7. Click **Save**."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("**3. Configure the Function App**\n\n1. In the Function App, select the Function App Name and select **Configuration**.\n2. In the **Application settings** tab, select **+ New application setting**.\n3. Add each of the following seven (7) application settings individually, with their respective string values (case-sensitive): \n		apiUsername\n		apiPassword\n		workspaceID\n		workspaceKey\n		uri\n		filterParameters\n		timeInterval\n> - Enter the URI that corresponds to your region. The complete list of API Server URLs can be [found here](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348). The `uri` value must follow the following schema: `https://<API Server>/api/2.0/fo/asset/host/vm/detection/?action=list&vm_processed_after=` -- There is no need to add a time suffix to the URI, the Function App will dynamically append the Time Value to the URI in the proper format.\n> - Add any additional filter parameters, for the `filterParameters` variable, that need to be appended to the URI. Each parameter should be seperated by an \"&\" symbol and should not include any spaces.\n> - Set the `timeInterval` (in minutes) to the value of `5` to correspond to the Timer Trigger of every `5` minutes. If the time interval needs to be modified, it is recommended to change the Function App Timer Trigger accordingly to prevent overlapping data ingestion.\n> - Note: If using Azure Key Vault, use the`@Microsoft.KeyVault(SecretUri={Security Identifier})`schema in place of the string values. Refer to [Key Vault references documentation](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references) for further details.\n4. Once all application settings have been entered, click **Save**."),
	// 							Title: to.Ptr(""),
	// 						},
	// 						{
	// 							Description: to.Ptr("**4. Configure the host.json**.\n\nDue to the potentially large amount of Qualys host detection data being ingested, it can cause the execution time to surpass the default Function App timeout of five (5) minutes. Increase the default timeout duration to the maximum of ten (10) minutes, under the Consumption Plan, to allow more time for the Function App to execute.\n\n1. In the Function App, select the Function App Name and select the **App Service Editor** blade.\n2. Click **Go** to open the editor, then select the **host.json** file under the **wwwroot** directory.\n3. Add the line `\"functionTimeout\": \"00:10:00\",` above the `managedDependancy` line \n4. Ensure **SAVED** appears on the top right corner of the editor, then exit the editor.\n\n> NOTE: If a longer timeout duration is required, consider upgrading to an [App Service Plan](https://docs.microsoft.com/azure/azure-functions/functions-scale#timeout)"),
	// 							Title: to.Ptr(""),
	// 					}},
	// 					Permissions: &armsecurityinsights.Permissions{
	// 						Customs: []*armsecurityinsights.PermissionsCustomsItem{
	// 							{
	// 								Name: to.Ptr("Microsoft.Web/sites permissions"),
	// 								Description: to.Ptr("Read and write permissions to Azure Functions to create a Function App is required. [See the documentation to learn more about Azure Functions](https://docs.microsoft.com/azure/azure-functions/)."),
	// 							},
	// 							{
	// 								Name: to.Ptr("Qualys API Key"),
	// 								Description: to.Ptr("A Qualys VM API username and password is required. [See the documentation to learn more about Qualys VM API](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf)."),
	// 						}},
	// 						ResourceProvider: []*armsecurityinsights.PermissionsResourceProviderItem{
	// 							{
	// 								PermissionsDisplayText: to.Ptr("read and write permissions on the workspace are required."),
	// 								Provider: to.Ptr(armsecurityinsights.ProviderNameMicrosoftOperationalInsightsWorkspaces),
	// 								ProviderDisplayName: to.Ptr("Workspace"),
	// 								RequiredPermissions: &armsecurityinsights.RequiredPermissions{
	// 									Delete: to.Ptr(true),
	// 									Read: to.Ptr(true),
	// 									Write: to.Ptr(true),
	// 								},
	// 								Scope: to.Ptr(armsecurityinsights.PermissionProviderScopeWorkspace),
	// 							},
	// 							{
	// 								PermissionsDisplayText: to.Ptr("read permissions to shared keys for the workspace are required. [See the documentation to learn more about workspace keys](https://docs.microsoft.com/azure/azure-monitor/platform/agent-windows#obtain-workspace-id-and-key)."),
	// 								Provider: to.Ptr(armsecurityinsights.ProviderNameMicrosoftOperationalInsightsWorkspacesSharedKeys),
	// 								ProviderDisplayName: to.Ptr("Keys"),
	// 								RequiredPermissions: &armsecurityinsights.RequiredPermissions{
	// 									Action: to.Ptr(true),
	// 								},
	// 								Scope: to.Ptr(armsecurityinsights.PermissionProviderScopeWorkspace),
	// 						}},
	// 					},
	// 					Publisher: to.Ptr("Qualys"),
	// 					SampleQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesSampleQueriesItem{
	// 						{
	// 							Description: to.Ptr("Top 10 Vulerabilities detected"),
	// 							Query: to.Ptr("{{graphQueriesTableName}}\n | mv-expand todynamic(Detections_s)\n | extend Vulnerability = tostring(Detections_s.Results)\n | summarize count() by Vulnerability\n | top 10 by count_"),
	// 					}},
	// 					Title: to.Ptr("Qualys Vulnerability Management (CCP DEMO)"),
	// 				},
	// 			},
	// 		},
	// 		                        }
}
