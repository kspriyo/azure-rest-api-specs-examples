package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/relations/GetEntityRelationByName.json
func ExampleEntityRelationsClient_GetRelation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntityRelationsClient().GetRelation(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Relation = armsecurityinsights.Relation{
	// 	Name: to.Ptr("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/entities/relations"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/afbd324f-6c48-459c-8710-8d1e1cd03812/relations/4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
	// 	Etag: to.Ptr("190057d0-0000-0d00-0000-5c6f5adb0000"),
	// 	Properties: &armsecurityinsights.RelationProperties{
	// 		RelatedResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/2216d0e1-91e3-4902-89fd-d2df8c535096"),
	// 		RelatedResourceName: to.Ptr("2216d0e1-91e3-4902-89fd-d2df8c535096"),
	// 		RelatedResourceType: to.Ptr("Microsoft.SecurityInsights/incidents"),
	// 	},
	// }
}
