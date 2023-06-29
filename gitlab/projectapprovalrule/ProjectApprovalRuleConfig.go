package projectapprovalrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectApprovalRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The number of approvals required for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#approvals_required ProjectApprovalRule#approvals_required}
	ApprovalsRequired *float64 `field:"required" json:"approvalsRequired" yaml:"approvalsRequired"`
	// The name of the approval rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#name ProjectApprovalRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or id of the project to add the approval rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#project ProjectApprovalRule#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// When this flag is set, the default `any_approver` rule will not be imported if present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#disable_importing_default_any_approver_rule_on_create ProjectApprovalRule#disable_importing_default_any_approver_rule_on_create}
	DisableImportingDefaultAnyApproverRuleOnCreate interface{} `field:"optional" json:"disableImportingDefaultAnyApproverRuleOnCreate" yaml:"disableImportingDefaultAnyApproverRuleOnCreate"`
	// A list of group IDs whose members can approve of the merge request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#group_ids ProjectApprovalRule#group_ids}
	GroupIds *[]*float64 `field:"optional" json:"groupIds" yaml:"groupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#id ProjectApprovalRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#protected_branch_ids ProjectApprovalRule#protected_branch_ids}
	ProtectedBranchIds *[]*float64 `field:"optional" json:"protectedBranchIds" yaml:"protectedBranchIds"`
	// String, defaults to 'regular'.
	//
	// The type of rule. `any_approver` is a pre-configured default rule with `approvals_required` at `0`. Valid values are `regular`, `any_approver`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#rule_type ProjectApprovalRule#rule_type}
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// A list of specific User IDs to add to the list of approvers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_approval_rule#user_ids ProjectApprovalRule#user_ids}
	UserIds *[]*float64 `field:"optional" json:"userIds" yaml:"userIds"`
}

