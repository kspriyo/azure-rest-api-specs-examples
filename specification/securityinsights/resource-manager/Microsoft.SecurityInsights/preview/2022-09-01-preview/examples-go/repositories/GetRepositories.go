package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/repositories/GetRepositories.json
func ExampleSourceControlClient_NewListRepositoriesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlClient().NewListRepositoriesPager("myRg", "myWorkspace", armsecurityinsights.RepoTypeGithub, nil)
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
		// page.RepoList = armsecurityinsights.RepoList{
		// 	Value: []*armsecurityinsights.Repo{
		// 		{
		// 			Branches: []*string{
		// 				to.Ptr("master"),
		// 				to.Ptr("develop")},
		// 				FullName: to.Ptr("reponame"),
		// 				URL: to.Ptr("https://api.github.com/repos/user/reponame"),
		// 		}},
		// 	}
	}
}
