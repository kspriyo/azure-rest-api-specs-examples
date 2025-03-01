package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListManagedInstanceOperations.json
func ExampleManagedInstanceOperationsClient_NewListByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceOperationsClient().NewListByManagedInstancePager("sqlcrudtest-7398", "sqlcrudtest-4645", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ManagedInstanceOperationListResult = armsql.ManagedInstanceOperationListResult{
		// 	Value: []*armsql.ManagedInstanceOperation{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/operations/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armsql.ManagedInstanceOperationProperties{
		// 				IsCancellable: to.Ptr(false),
		// 				ManagedInstanceName: to.Ptr("sqlcrudtest-4645"),
		// 				Operation: to.Ptr("UpsertManagedServer"),
		// 				OperationFriendlyName: to.Ptr("UPDATE MANAGED SERVER"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-06T11:08:44.490Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateCancelled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/operations/22222222-2222-2222-2222-222222222222"),
		// 			Properties: &armsql.ManagedInstanceOperationProperties{
		// 				IsCancellable: to.Ptr(true),
		// 				ManagedInstanceName: to.Ptr("sqlcrudtest-4645"),
		// 				Operation: to.Ptr("UpsertManagedServer"),
		// 				OperationFriendlyName: to.Ptr("UPDATE MANAGED SERVER"),
		// 				OperationParameters: &armsql.ManagedInstanceOperationParametersPair{
		// 					CurrentParameters: &armsql.UpsertManagedServerOperationParameters{
		// 						Family: to.Ptr("Gen4"),
		// 						StorageSizeInGB: to.Ptr[int32](32),
		// 						Tier: to.Ptr("GeneralPurpose"),
		// 						VCores: to.Ptr[int32](8),
		// 					},
		// 					RequestedParameters: &armsql.UpsertManagedServerOperationParameters{
		// 						Family: to.Ptr("Gen4"),
		// 						StorageSizeInGB: to.Ptr[int32](128),
		// 						Tier: to.Ptr("BusinessCritical"),
		// 						VCores: to.Ptr[int32](8),
		// 					},
		// 				},
		// 				OperationSteps: &armsql.ManagedInstanceOperationSteps{
		// 					CurrentStep: to.Ptr[int32](3),
		// 					StepsList: []*armsql.UpsertManagedServerOperationStep{
		// 						{
		// 							Name: to.Ptr("Request validation"),
		// 							Order: to.Ptr[int32](1),
		// 							Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusCompleted),
		// 						},
		// 						{
		// 							Name: to.Ptr("Virtual Cluster resize/creation"),
		// 							Order: to.Ptr[int32](2),
		// 							Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusCompleted),
		// 						},
		// 						{
		// 							Name: to.Ptr("New SQL Instance Startup"),
		// 							Order: to.Ptr[int32](3),
		// 							Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusInProgress),
		// 					}},
		// 					TotalSteps: to.Ptr("3"),
		// 				},
		// 				PercentComplete: to.Ptr[int32](50),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-06T11:08:44.490Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateInProgress),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/operations/33333333-3333-3333-3333-333333333333"),
		// 			Properties: &armsql.ManagedInstanceOperationProperties{
		// 				ErrorCode: to.Ptr[int32](45157),
		// 				ErrorDescription: to.Ptr("Server 'sqlcrudtest-4645' is busy with another operation. Please try your operation later."),
		// 				ErrorSeverity: to.Ptr[int32](16),
		// 				IsCancellable: to.Ptr(false),
		// 				IsUserError: to.Ptr(true),
		// 				ManagedInstanceName: to.Ptr("sqlcrudtest-4645"),
		// 				Operation: to.Ptr("UpsertManagedServer"),
		// 				OperationFriendlyName: to.Ptr("UPDATE MANAGED SERVER"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-06T14:48:34.583Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateFailed),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("44444444-4444-4444-4444-444444444444"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/operations/44444444-4444-4444-4444-444444444444"),
		// 			Properties: &armsql.ManagedInstanceOperationProperties{
		// 				ManagedInstanceName: to.Ptr("sqlcrudtest-4645"),
		// 				Operation: to.Ptr("MakeAllManagedDatabasesAccessible"),
		// 				OperationFriendlyName: to.Ptr("MAKE ALL MANAGED DBS ACCESSIBLE"),
		// 				OperationSteps: &armsql.ManagedInstanceOperationSteps{
		// 					CurrentStep: to.Ptr[int32](1),
		// 					StepsList: []*armsql.UpsertManagedServerOperationStep{
		// 						{
		// 							Name: to.Ptr("Request validation"),
		// 							Order: to.Ptr[int32](1),
		// 							Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusFailed),
		// 					}},
		// 					TotalSteps: to.Ptr("1"),
		// 				},
		// 				PercentComplete: to.Ptr[int32](100),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T09:00:08.100Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
