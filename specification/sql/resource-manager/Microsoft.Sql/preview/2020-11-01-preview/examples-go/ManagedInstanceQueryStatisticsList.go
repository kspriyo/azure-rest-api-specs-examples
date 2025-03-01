package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceQueryStatisticsList.json
func ExampleManagedDatabaseQueriesClient_NewListByQueryPager_obtainQueryExecutionStatistics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseQueriesClient().NewListByQueryPager("sqlcrudtest-7398", "sqlcrudtest-4645", "database_1", "42", &armsql.ManagedDatabaseQueriesClientListByQueryOptions{StartTime: nil,
		EndTime:  nil,
		Interval: nil,
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
		// page.ManagedInstanceQueryStatistics = armsql.ManagedInstanceQueryStatistics{
		// 	Value: []*armsql.QueryStatistics{
		// 		{
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/queries/statistics"),
		// 			ID: to.Ptr("35"),
		// 			Properties: &armsql.QueryStatisticsProperties{
		// 				DatabaseName: to.Ptr("db1"),
		// 				EndTime: to.Ptr("03/02/2020 18:34:58"),
		// 				Intervals: []*armsql.QueryMetricInterval{
		// 					{
		// 						ExecutionCount: to.Ptr[int64](160),
		// 						IntervalStartTime: to.Ptr("03/02/2020 08:00:00"),
		// 						IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 						Metrics: []*armsql.QueryMetricProperties{
		// 							{
		// 								Name: to.Ptr("cpu"),
		// 								Avg: to.Ptr[float64](0.00001665347222222222),
		// 								DisplayName: to.Ptr("Cpu"),
		// 								Max: to.Ptr[float64](0.000025243055555555557),
		// 								Min: to.Ptr[float64](0.00001507638888888889),
		// 								Stdev: to.Ptr[float64](0.0000014894345929850385),
		// 								Sum: to.Ptr[float64](0.0026645555555555554),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("io"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Physical Io Reads"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("logIo"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Log Writes"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("memory"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Memory consumption"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("duration"),
		// 								Avg: to.Ptr[float64](5026.625),
		// 								DisplayName: to.Ptr("Query duration"),
		// 								Max: to.Ptr[float64](18490),
		// 								Min: to.Ptr[float64](4373),
		// 								Stdev: to.Ptr[float64](1487.3520882343225),
		// 								Sum: to.Ptr[float64](804260),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 								Value: to.Ptr[float64](0),
		// 						}},
		// 					},
		// 					{
		// 						ExecutionCount: to.Ptr[int64](20),
		// 						IntervalStartTime: to.Ptr("03/02/2020 09:00:00"),
		// 						IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 						Metrics: []*armsql.QueryMetricProperties{
		// 							{
		// 								Name: to.Ptr("cpu"),
		// 								Avg: to.Ptr[float64](0.00004479774305555555),
		// 								DisplayName: to.Ptr("Cpu"),
		// 								Max: to.Ptr[float64](0.00014645833333333332),
		// 								Min: to.Ptr[float64](0.000023430555555555557),
		// 								Stdev: to.Ptr[float64](0.00003830118344204395),
		// 								Sum: to.Ptr[float64](0.0008959548611111111),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("io"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Physical Io Reads"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("logIo"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Log Writes"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("memory"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Memory consumption"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("duration"),
		// 								Avg: to.Ptr[float64](12963.2),
		// 								DisplayName: to.Ptr("Query duration"),
		// 								Max: to.Ptr[float64](42289),
		// 								Min: to.Ptr[float64](6813),
		// 								Stdev: to.Ptr[float64](11040.140794392071),
		// 								Sum: to.Ptr[float64](259264),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 								Value: to.Ptr[float64](0),
		// 						}},
		// 					},
		// 					{
		// 						ExecutionCount: to.Ptr[int64](80),
		// 						IntervalStartTime: to.Ptr("03/02/2020 15:00:00"),
		// 						IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 						Metrics: []*armsql.QueryMetricProperties{
		// 							{
		// 								Name: to.Ptr("cpu"),
		// 								Avg: to.Ptr[float64](0.000019315538194444445),
		// 								DisplayName: to.Ptr("Cpu"),
		// 								Max: to.Ptr[float64](0.00002764236111111111),
		// 								Min: to.Ptr[float64](0.000018215277777777777),
		// 								Stdev: to.Ptr[float64](0.0000010716305801875179),
		// 								Sum: to.Ptr[float64](0.0015452430555555556),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("io"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Physical Io Reads"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("logIo"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Log Writes"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("memory"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Memory consumption"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("duration"),
		// 								Avg: to.Ptr[float64](5586.2625),
		// 								DisplayName: to.Ptr("Query duration"),
		// 								Max: to.Ptr[float64](7982),
		// 								Min: to.Ptr[float64](5264),
		// 								Stdev: to.Ptr[float64](310.4915757210727),
		// 								Sum: to.Ptr[float64](446901),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 								Value: to.Ptr[float64](0),
		// 						}},
		// 					},
		// 					{
		// 						ExecutionCount: to.Ptr[int64](80),
		// 						IntervalStartTime: to.Ptr("03/02/2020 17:00:00"),
		// 						IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 						Metrics: []*armsql.QueryMetricProperties{
		// 							{
		// 								Name: to.Ptr("cpu"),
		// 								Avg: to.Ptr[float64](0.000019085373263888888),
		// 								DisplayName: to.Ptr("Cpu"),
		// 								Max: to.Ptr[float64](0.00002782638888888889),
		// 								Min: to.Ptr[float64](0.000017819444444444443),
		// 								Stdev: to.Ptr[float64](0.0000012309244108727927),
		// 								Sum: to.Ptr[float64](0.0015268298611111112),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("io"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Physical Io Reads"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("logIo"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Log Writes"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("memory"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Memory consumption"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("duration"),
		// 								Avg: to.Ptr[float64](5517.2),
		// 								DisplayName: to.Ptr("Query duration"),
		// 								Max: to.Ptr[float64](8052),
		// 								Min: to.Ptr[float64](5147),
		// 								Stdev: to.Ptr[float64](356.8763581410226),
		// 								Sum: to.Ptr[float64](441376),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 								Value: to.Ptr[float64](0),
		// 						}},
		// 				}},
		// 				QueryID: to.Ptr("35"),
		// 				StartTime: to.Ptr("03/01/2020 18:34:58"),
		// 			},
		// 	}},
		// }
	}
}
