package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/CreateFileImport.json
func ExampleFileImportsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewFileImportsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.FileImport{
		Properties: &armsecurityinsights.FileImportProperties{
			ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
			ImportFile: &armsecurityinsights.FileMetadata{
				FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
				FileName:   to.Ptr("myFile.json"),
				FileSize:   to.Ptr[int32](4653),
			},
			IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
			Source:        to.Ptr("mySource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
