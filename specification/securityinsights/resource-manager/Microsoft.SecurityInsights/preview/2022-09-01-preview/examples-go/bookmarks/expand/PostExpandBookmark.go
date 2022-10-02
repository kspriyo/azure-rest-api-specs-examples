package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/expand/PostExpandBookmark.json
func ExampleBookmarkClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewBookmarkClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Expand(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.BookmarkExpandParameters{
		EndTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-24T17:21:00.000Z"); return t }()),
		ExpansionID: to.Ptr("27f76e63-c41b-480f-bb18-12ad2e011d49"),
		StartTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-25T17:21:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
