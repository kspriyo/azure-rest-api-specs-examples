package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionTargetsByExecution.json
func ExampleJobTargetExecutionsClient_NewListByJobExecutionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobTargetExecutionsClient().NewListByJobExecutionPager("group1", "server1", "agent1", "job1", "5A86BF65-43AC-F258-2524-9E92992F97CA", &armsql.JobTargetExecutionsClientListByJobExecutionOptions{CreateTimeMin: nil,
		CreateTimeMax: nil,
		EndTimeMin:    nil,
		EndTimeMax:    nil,
		IsActive:      nil,
		Skip:          nil,
		Top:           nil,
	})
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
		// page.JobExecutionListResult = armsql.JobExecutionListResult{
		// 	Value: []*armsql.JobExecution{
		// 		{
		// 			Name: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/executions/steps/targets"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/executions/5555-6666-7777-8888-999999999999/steps/step1/targets/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 			Properties: &armsql.JobExecutionProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T04:33:17.513Z"); return t}()),
		// 				CurrentAttemptStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T04:33:18.239Z"); return t}()),
		// 				CurrentAttempts: to.Ptr[int32](1),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T04:33:18.703Z"); return t}()),
		// 				JobExecutionID: to.Ptr("5A86BF65-43AC-F258-2524-9E92992F97CA"),
		// 				JobVersion: to.Ptr[int32](1),
		// 				LastMessage: to.Ptr("Step 1 succeeded execution on target (server 'server1', database 'database1')."),
		// 				Lifecycle: to.Ptr(armsql.JobExecutionLifecycleSucceeded),
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T04:33:18.123Z"); return t}()),
		// 				StepID: to.Ptr[int32](1),
		// 				StepName: to.Ptr("step1"),
		// 				Target: &armsql.JobExecutionTarget{
		// 					Type: to.Ptr(armsql.JobTargetTypeSQLDatabase),
		// 					DatabaseName: to.Ptr("database1"),
		// 					ServerName: to.Ptr("server1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
