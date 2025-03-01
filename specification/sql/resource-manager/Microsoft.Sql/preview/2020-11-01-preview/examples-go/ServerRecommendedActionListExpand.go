package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerRecommendedActionListExpand.json
func ExampleServerAdvisorsClient_ListByServer_listOfServerRecommendedActionsForAllAdvisors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerAdvisorsClient().ListByServer(ctx, "workloadinsight-demos", "misosisvr", &armsql.ServerAdvisorsClientListByServerOptions{Expand: to.Ptr("recommendedActions")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdvisorArray = []*armsql.Advisor{
	// 	{
	// 		Name: to.Ptr("CreateIndex"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/CreateIndex"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromServer),
	// 			RecommendedActions: []*armsql.RecommendedAction{
	// 				{
	// 					Name: to.Ptr("IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/CreateIndex/recommendedActions/IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](1440),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](209.3125),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B] ON [CRM].[DataPoints] ([Name],[Money],[Power]) INCLUDE ([Hour], [System], [LastChanged]) WITH (ONLINE = ON)"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr(""),
	// 						Score: to.Ptr[int32](1),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T15:15:40.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[CRM]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Hour], [System], [LastChanged]",
	// 							"indexColumns": "[Name],[Money],[Power]",
	// 							"indexName": "nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B",
	// 							"indexType": "NONCLUSTERED",
	// 							"table": "[DataPoints]",
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("IR_[dbo]_[Employees]_560E15A98D14CA09BDFB"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/CreateIndex/recommendedActions/IR_[dbo]_[Employees]_560E15A98D14CA09BDFB"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](17),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](128),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_Employees_8C18C2AF4267DC77793040782641CCDE] ON [dbo].[Employees] ([City], [State]) INCLUDE ([Postal]) WITH (ONLINE = ON)"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr(""),
	// 						Score: to.Ptr[int32](3),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[dbo]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Postal]",
	// 							"indexColumns": "[City], [State]",
	// 							"indexName": "nci_wi_Employees_8C18C2AF4267DC77793040782641CCDE",
	// 							"indexType": "NONCLUSTERED",
	// 							"table": "[Employees]",
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("IR_[dbo]_[DataPoints]_F5D2F347AA22DB46E4CC"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/CreateIndex/recommendedActions/IR_[dbo]_[DataPoints]_F5D2F347AA22DB46E4CC"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](5040),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](120),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ExecuteActionDuration: to.Ptr("PT1M"),
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ExecuteActionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("DROP INDEX [nci_wi_DataPoints_609E4B7D6A3813990ED44B28B340C8FC] ON [dbo].[DataPoints]"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](-12.7),
	// 								ChangeValueRelative: to.Ptr[float64](-0.9),
	// 								DimensionName: to.Ptr("AffectedQueriesCpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](-12.7),
	// 								ChangeValueRelative: to.Ptr[float64](-0.3175),
	// 								DimensionName: to.Ptr("CpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](12),
	// 								DimensionName: to.Ptr("QueriesWithImprovedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](1),
	// 								DimensionName: to.Ptr("QueriesWithRegressedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](130.742187),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("VerificationProgress"),
	// 								Unit: to.Ptr("Percent"),
	// 						}},
	// 						RecommendationReason: to.Ptr(""),
	// 						Score: to.Ptr[int32](3),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateSuccess),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[dbo]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Power],[Pineapple]",
	// 							"indexActionDuration": "00:01:00",
	// 							"indexActionStartTime": "2017-03-01T14:38:05.337",
	// 							"indexColumns": "[Name],[Money]",
	// 							"indexName": "nci_wi_DataPoints_609E4B7D6A3813990ED44B28B340C8FC",
	// 							"indexType": "NONCLUSTERED",
	// 							"table": "[DataPoints]",
	// 						},
	// 					},
	// 			}},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("DropIndex"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 			RecommendedActions: []*armsql.RecommendedAction{
	// 				{
	// 					Name: to.Ptr("IR_[CRM]_[DataPoints1]_29AEA82685D24704DE1A"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex/recommendedActions/IR_[CRM]_[DataPoints1]_29AEA82685D24704DE1A"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](803),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](144.6875),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("DROP INDEX [MyIndex123] ON [CRM].[DataPoints1]"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr("Duplicate"),
	// 						Score: to.Ptr[int32](1),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[CRM]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Apple]",
	// 							"indexColumns": "[Cookies],[SessionId]",
	// 							"indexName": "MyIndex123",
	// 							"indexType": "NONCLUSTERED",
	// 							"originalIndexName": "IX_COM_SKU_SKUDepartmentID",
	// 							"table": "[DataPoints1]",
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("IR_[CRM]_[DataPoints2]_E4B21F229379807E531A"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex/recommendedActions/IR_[CRM]_[DataPoints2]_E4B21F229379807E531A"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](705),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](342),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ExecuteActionDuration: to.Ptr("PT1M"),
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ExecuteActionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("CREATE NONCLUSTERED INDEX [MyIndex321] ON [CRM].[DataPoints2] ([Cookies],[SessionId],[Protocol]) INCLUDE ([Apple]) WITH (ONLINE = ON)"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](0),
	// 								ChangeValueRelative: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("AffectedQueriesCpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](0),
	// 								ChangeValueRelative: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("CpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("QueriesWithImprovedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("QueriesWithRegressedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](-342),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("VerificationProgress"),
	// 								Unit: to.Ptr("Percent"),
	// 						}},
	// 						RecommendationReason: to.Ptr("Duplicate"),
	// 						Score: to.Ptr[int32](1),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateSuccess),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[CRM]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Apple]",
	// 							"indexActionDuration": "00:01:00",
	// 							"indexActionStartTime": "2017-03-01T14:38:05.697",
	// 							"indexColumns": "[Cookies],[SessionId],[Protocol]",
	// 							"indexName": "MyIndex321",
	// 							"indexType": "NONCLUSTERED",
	// 							"originalIndexName": "IX_COM_SKU_SKUDepartmentID",
	// 							"table": "[DataPoints2]",
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("IR_[dbo]_[FactFinance]_13ADA5F354E9E14A983B"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex/recommendedActions/IR_[dbo]_[FactFinance]_13ADA5F354E9E14A983B"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](705),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](342),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ExecuteActionDuration: to.Ptr("PT1M"),
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ExecuteActionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("CREATE NONCLUSTERED INDEX [IX_FF] ON [dbo].[FactFinance] ([FinanceKey],[DateKey]) INCLUDE ([OrganizationKey]) WITH (ONLINE = ON)"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](0),
	// 								ChangeValueRelative: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("AffectedQueriesCpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](0),
	// 								ChangeValueRelative: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("CpuUtilization"),
	// 								Unit: to.Ptr("CpuCores"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("QueriesWithImprovedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("QueriesWithRegressedPerformance"),
	// 								Unit: to.Ptr("Count"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](-342),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](0),
	// 								DimensionName: to.Ptr("VerificationProgress"),
	// 								Unit: to.Ptr("Percent"),
	// 						}},
	// 						RecommendationReason: to.Ptr("Duplicate"),
	// 						Score: to.Ptr[int32](1),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateSuccess),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-28T14:38:05.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[dbo]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[OrganizationKey]",
	// 							"indexActionDuration": "00:01:00",
	// 							"indexActionStartTime": "2017-03-01T14:38:05.837",
	// 							"indexColumns": "[FinanceKey],[DateKey]",
	// 							"indexName": "IX_FF",
	// 							"indexType": "NONCLUSTERED",
	// 							"originalIndexName": "IX_COM_SKU_SKUDepartmentID",
	// 							"table": "[FactFinance]",
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("IR_[CRM]_[DataPoints1]_13ADA5F354E9E14A983B"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex/recommendedActions/IR_[CRM]_[DataPoints1]_13ADA5F354E9E14A983B"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](645),
	// 								DimensionName: to.Ptr("ActionDuration"),
	// 								Unit: to.Ptr("Seconds"),
	// 							},
	// 							{
	// 								AbsoluteValue: to.Ptr[float64](342),
	// 								DimensionName: to.Ptr("SpaceChange"),
	// 								Unit: to.Ptr("Megabytes"),
	// 						}},
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-23T14:14:35.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("DROP INDEX [IX_FF] ON [CRM].[DataPoints1]"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr("Unused"),
	// 						Score: to.Ptr[int32](1),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-23T14:14:35.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"schema": "[CRM]",
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"includedColumns": "[Apple]",
	// 							"indexColumns": "[Protocol],[SessionId]",
	// 							"indexName": "IX_FF",
	// 							"indexType": "NONCLUSTERED",
	// 							"originalIndexName": "IX_COM_SKU_SKUDepartmentID",
	// 							"table": "[DataPoints1]",
	// 						},
	// 					},
	// 			}},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("DbParameterization"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DbParameterization"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 			RecommendedActions: []*armsql.RecommendedAction{
	// 				{
	// 					Name: to.Ptr("ForceDbParameterization"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DbParameterization/recommendedActions/ForceDbParameterization"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](22.5613696939135),
	// 								DimensionName: to.Ptr("CpuSavings"),
	// 								Unit: to.Ptr("Percent"),
	// 							},
	// 							{
	// 								ChangeValueAbsolute: to.Ptr[float64](0.701823681806341),
	// 								DimensionName: to.Ptr("QueryDurationDecrease"),
	// 								Unit: to.Ptr("secs"),
	// 						}},
	// 						ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 						ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T15:08:31.000Z"); return t}()),
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 							Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 							Script: to.Ptr("ALTER DATABASE [IndexAdvisor_test_3] SET PARAMETERIZATION FORCED"),
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(true),
	// 						IsRevertableAction: to.Ptr(true),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:37:44.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr(""),
	// 						Score: to.Ptr[int32](3),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedBySystem),
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T15:08:31.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:37:44.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"databaseName": "IndexAdvisor_test_3",
	// 						},
	// 					},
	// 			}},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("SchemaIssue"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/SchemaIssue"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusPublicPreview),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 			RecommendedActions: []*armsql.RecommendedAction{
	// 				{
	// 					Name: to.Ptr("SchemaProblem_1A258C5714A7410C9D23"),
	// 					Type: to.Ptr("Microsoft.Sql/servers/advisors/recommendedActions"),
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/SchemaIssue/recommendedActions/SchemaProblem_1A258C5714A7410C9D23"),
	// 					Kind: to.Ptr(""),
	// 					Location: to.Ptr("East Asia"),
	// 					Properties: &armsql.RecommendedActionProperties{
	// 						ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 						},
	// 						EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 						},
	// 						IsArchivedAction: to.Ptr(false),
	// 						IsExecutableAction: to.Ptr(false),
	// 						IsRevertableAction: to.Ptr(false),
	// 						LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T22:19:53.000Z"); return t}()),
	// 						ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 						},
	// 						RecommendationReason: to.Ptr("SchemaProblem"),
	// 						Score: to.Ptr[int32](3),
	// 						State: &armsql.RecommendedActionStateInfo{
	// 							CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 							LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:37:48.000Z"); return t}()),
	// 						},
	// 						TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 						},
	// 						ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:37:48.000Z"); return t}()),
	// 						Details: map[string]any{
	// 							"databaseName": "IndexAdvisor_test_3",
	// 							"sqlErrorCount": float64(342482),
	// 							"sqlErrorMessage": "Invalid object name 'dbo.Companies'.",
	// 							"sqlErrorNumber": float64(208),
	// 						},
	// 					},
	// 			}},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("ForceLastGoodPlan"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/ForceLastGoodPlan"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusPrivatePreview),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 			RecommendedActions: []*armsql.RecommendedAction{
	// 			},
	// 		},
	// }}
}
