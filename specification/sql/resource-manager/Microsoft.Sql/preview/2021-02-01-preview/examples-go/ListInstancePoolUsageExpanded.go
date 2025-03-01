package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListInstancePoolUsageExpanded.json
func ExampleUsagesClient_NewListByInstancePoolPager_listInstancePoolUsagesExpandedWithChildren() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListByInstancePoolPager("group1", "testIP", &armsql.UsagesClientListByInstancePoolOptions{ExpandChildren: to.Ptr(true)})
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
		// page.UsageListResult = armsql.UsageListResult{
		// 	Value: []*armsql.Usage{
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("VCore utilization"),
		// 				Value: to.Ptr("VCore utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/usages"),
		// 			CurrentValue: to.Ptr[int32](12),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/usages/vcore_utilization"),
		// 			Limit: to.Ptr[int32](16),
		// 			RequestedLimit: to.Ptr[int32](40),
		// 			Unit: to.Ptr("VCores"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("Storage utilization"),
		// 				Value: to.Ptr("Storage utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/usages"),
		// 			CurrentValue: to.Ptr[int32](384),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/usages/storage_utilization"),
		// 			Limit: to.Ptr[int32](8196),
		// 			Unit: to.Ptr("VCores"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("Database utilization"),
		// 				Value: to.Ptr("Database utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/usages"),
		// 			CurrentValue: to.Ptr[int32](5),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/usages/database_utilization"),
		// 			Limit: to.Ptr[int32](100),
		// 			Unit: to.Ptr("Number Of Databases"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("VCore utilization"),
		// 				Value: to.Ptr("VCore utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance1/usages/vcore_utilization"),
		// 			Limit: to.Ptr[int32](4),
		// 			Unit: to.Ptr("VCores"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("VCore utilization"),
		// 				Value: to.Ptr("VCore utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance2/usages/vcore_utilization"),
		// 			Limit: to.Ptr[int32](4),
		// 			RequestedLimit: to.Ptr[int32](8),
		// 			Unit: to.Ptr("VCores"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("Storage utilization"),
		// 				Value: to.Ptr("Storage utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance1/usages/storage_utilization"),
		// 			Limit: to.Ptr[int32](128),
		// 			Unit: to.Ptr("Gigabytes"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("VCore utilization"),
		// 				Value: to.Ptr("VCore utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance2/usages/storage_utilization"),
		// 			Limit: to.Ptr[int32](128),
		// 			RequestedLimit: to.Ptr[int32](256),
		// 			Unit: to.Ptr("Gigabytes"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("Database utilization"),
		// 				Value: to.Ptr("Database utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			CurrentValue: to.Ptr[int32](2),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance1/usages/database_utilization"),
		// 			Limit: to.Ptr[int32](100),
		// 			Unit: to.Ptr("Number Of Databases"),
		// 		},
		// 		{
		// 			Name: &armsql.Name{
		// 				LocalizedValue: to.Ptr("Database utilization"),
		// 				Value: to.Ptr("Database utilization"),
		// 			},
		// 			Type: to.Ptr("Microsoft.Sql/instancePools/managedInstances/usages"),
		// 			CurrentValue: to.Ptr[int32](3),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP/managedInstances/managedInstance2/usages/database_utilization"),
		// 			Limit: to.Ptr[int32](100),
		// 			Unit: to.Ptr("Number Of Databases"),
		// 	}},
		// }
	}
}
