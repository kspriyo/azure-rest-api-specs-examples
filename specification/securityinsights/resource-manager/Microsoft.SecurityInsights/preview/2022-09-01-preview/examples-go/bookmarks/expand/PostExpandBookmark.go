package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/expand/PostExpandBookmark.json
func ExampleBookmarkClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBookmarkClient().Expand(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.BookmarkExpandParameters{
		EndTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-24T17:21:00.000Z"); return t }()),
		ExpansionID: to.Ptr("27f76e63-c41b-480f-bb18-12ad2e011d49"),
		StartTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-25T17:21:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BookmarkExpandResponse = armsecurityinsights.BookmarkExpandResponse{
	// 	MetaData: &armsecurityinsights.ExpansionResultsMetadata{
	// 		Aggregations: []*armsecurityinsights.ExpansionResultAggregation{
	// 			{
	// 				Count: to.Ptr[int32](1),
	// 				EntityKind: to.Ptr(armsecurityinsights.EntityKindAccount),
	// 		}},
	// 	},
	// 	Value: &armsecurityinsights.BookmarkExpandResponseValue{
	// 		Entities: []armsecurityinsights.EntityClassification{
	// 			&armsecurityinsights.AccountEntity{
	// 				Name: to.Ptr("fe4ddab5-8cea-eca3-c8b8-9e92e830a387"),
	// 				Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/fe4ddab5-8cea-eca3-c8b8-9e92e830a387"),
	// 				Kind: to.Ptr(armsecurityinsights.EntityKindAccount),
	// 				Properties: &armsecurityinsights.AccountEntityProperties{
	// 					FriendlyName: to.Ptr("administrator"),
	// 					AccountName: to.Ptr("administrator"),
	// 					NtDomain: to.Ptr("domain"),
	// 				},
	// 		}},
	// 	},
	// }
}
