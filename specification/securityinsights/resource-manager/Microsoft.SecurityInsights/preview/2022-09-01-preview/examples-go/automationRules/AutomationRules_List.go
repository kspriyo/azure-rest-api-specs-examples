package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/automationRules/AutomationRules_List.json
func ExampleAutomationRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutomationRulesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.AutomationRulesList = armsecurityinsights.AutomationRulesList{
		// 	Value: []*armsecurityinsights.AutomationRule{
		// 		{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/automationRules"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/automationRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.AutomationRuleProperties{
		// 				Actions: []armsecurityinsights.AutomationRuleActionClassification{
		// 					&armsecurityinsights.AutomationRuleRunPlaybookAction{
		// 						ActionType: to.Ptr(armsecurityinsights.ActionTypeRunPlaybook),
		// 						Order: to.Ptr[int32](1),
		// 						ActionConfiguration: &armsecurityinsights.PlaybookActionProperties{
		// 							LogicAppResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/AlertPlaybook"),
		// 							TenantID: to.Ptr("d23e3eef-eed0-428f-a2d5-bc48c268e31d"),
		// 						},
		// 				}},
		// 				CreatedBy: &armsecurityinsights.ClientInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john.doe@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 					UserPrincipalName: to.Ptr("john@contoso.com"),
		// 				},
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:00Z"); return t}()),
		// 				DisplayName: to.Ptr("Suspicious alerts in workspace"),
		// 				LastModifiedBy: &armsecurityinsights.ClientInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john.doe@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 					UserPrincipalName: to.Ptr("john@contoso.com"),
		// 				},
		// 				LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t}()),
		// 				Order: to.Ptr[int32](1),
		// 				TriggeringLogic: &armsecurityinsights.AutomationRuleTriggeringLogic{
		// 					Conditions: []armsecurityinsights.AutomationRuleConditionClassification{
		// 						&armsecurityinsights.PropertyConditionProperties{
		// 							ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
		// 							ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
		// 								Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorContains),
		// 								PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyAlertAnalyticRuleIDs),
		// 								PropertyValues: []*string{
		// 									to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/fab3d2d4-747f-46a7-8ef0-9c0be8112bf7"),
		// 									to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/8deb8303-e94d-46ff-96e0-5fd94b33df1a")},
		// 								},
		// 						}},
		// 						IsEnabled: to.Ptr(true),
		// 						TriggersOn: to.Ptr(armsecurityinsights.TriggersOnAlerts),
		// 						TriggersWhen: to.Ptr(armsecurityinsights.TriggersWhenCreated),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
