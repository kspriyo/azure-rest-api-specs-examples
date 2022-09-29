package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PatchRoleManagementPolicy.json
func ExampleRoleManagementPoliciesClient_Update_patchRoleManagementPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleManagementPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", "570c3619-7688-4b34-b290-2b8bb3ccab2a", armauthorization.RoleManagementPolicy{
		Properties: &armauthorization.RoleManagementPolicyProperties{
			Rules: []armauthorization.RoleManagementPolicyRuleClassification{
				&armauthorization.RoleManagementPolicyExpirationRule{
					ID:       to.Ptr("Expiration_Admin_Eligibility"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsExpirationRequired: to.Ptr(false),
					MaximumDuration:      to.Ptr("P180D"),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Admin_Admin_Eligibility"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("admin_admin_eligible@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeAdmin),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Requestor_Admin_Eligibility"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("requestor_admin_eligible@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeRequestor),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Approver_Admin_Eligibility"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("approver_admin_eligible@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeApprover),
				},
				&armauthorization.RoleManagementPolicyEnablementRule{
					ID:       to.Ptr("Enablement_Admin_Eligibility"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All")},
					},
					EnabledRules: []*armauthorization.EnablementRules{},
				},
				&armauthorization.RoleManagementPolicyExpirationRule{
					ID:       to.Ptr("Expiration_Admin_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsExpirationRequired: to.Ptr(false),
					MaximumDuration:      to.Ptr("P90D"),
				},
				&armauthorization.RoleManagementPolicyEnablementRule{
					ID:       to.Ptr("Enablement_Admin_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					EnabledRules: []*armauthorization.EnablementRules{
						to.Ptr(armauthorization.EnablementRulesJustification),
						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication)},
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Admin_Admin_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("admin_admin_member@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeAdmin),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Requestor_Admin_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("requestor_admin_member@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeRequestor),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Approver_Admin_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("approver_admin_member@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeApprover),
				},
				&armauthorization.RoleManagementPolicyExpirationRule{
					ID:       to.Ptr("Expiration_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsExpirationRequired: to.Ptr(true),
					MaximumDuration:      to.Ptr("PT7H"),
				},
				&armauthorization.RoleManagementPolicyEnablementRule{
					ID:       to.Ptr("Enablement_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					EnabledRules: []*armauthorization.EnablementRules{
						to.Ptr(armauthorization.EnablementRulesJustification),
						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
						to.Ptr(armauthorization.EnablementRulesTicketing)},
				},
				&armauthorization.RoleManagementPolicyApprovalRule{
					ID:       to.Ptr("Approval_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					Setting: &armauthorization.ApprovalSettings{
						ApprovalMode: to.Ptr(armauthorization.ApprovalModeSingleStage),
						ApprovalStages: []*armauthorization.ApprovalStage{
							{
								ApprovalStageTimeOutInDays:      to.Ptr[int32](1),
								EscalationTimeInMinutes:         to.Ptr[int32](0),
								IsApproverJustificationRequired: to.Ptr(true),
								IsEscalationEnabled:             to.Ptr(false),
								PrimaryApprovers: []*armauthorization.UserSet{
									{
										Description: to.Ptr("amansw_new_group"),
										ID:          to.Ptr("2385b0f3-5fa9-43cf-8ca4-b01dc97298cd"),
										IsBackup:    to.Ptr(false),
										UserType:    to.Ptr(armauthorization.UserTypeGroup),
									},
									{
										Description: to.Ptr("amansw_group"),
										ID:          to.Ptr("2f4913c9-d15b-406a-9946-1d66a28f2690"),
										IsBackup:    to.Ptr(false),
										UserType:    to.Ptr(armauthorization.UserTypeGroup),
									}},
							}},
						IsApprovalRequired:               to.Ptr(true),
						IsApprovalRequiredForExtension:   to.Ptr(false),
						IsRequestorJustificationRequired: to.Ptr(true),
					},
				},
				&armauthorization.RoleManagementPolicyAuthenticationContextRule{
					ID:       to.Ptr("AuthenticationContext_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					ClaimValue: to.Ptr(""),
					IsEnabled:  to.Ptr(false),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Admin_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("admin_enduser_member@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeAdmin),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Requestor_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("requestor_enduser_member@test.com")},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeRequestor),
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:       to.Ptr("Notification_Approver_EndUser_Assignment"),
					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("EndUser"),
						Level:  to.Ptr("Assignment"),
						Operations: []*string{
							to.Ptr("All")},
					},
					IsDefaultRecipientsEnabled: to.Ptr(true),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationType:           to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:              to.Ptr(armauthorization.RecipientTypeApprover),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
