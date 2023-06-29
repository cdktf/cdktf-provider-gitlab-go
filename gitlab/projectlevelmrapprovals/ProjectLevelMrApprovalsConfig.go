package projectlevelmrapprovals

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectLevelMrApprovalsConfig struct {
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
	// The ID or URL-encoded path of a project to change MR approval configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#project ProjectLevelMrApprovals#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// By default, users are able to edit the approval rules in merge requests. If set to true,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#disable_overriding_approvers_per_merge_request ProjectLevelMrApprovals#disable_overriding_approvers_per_merge_request}
	DisableOverridingApproversPerMergeRequest interface{} `field:"optional" json:"disableOverridingApproversPerMergeRequest" yaml:"disableOverridingApproversPerMergeRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#id ProjectLevelMrApprovals#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#merge_requests_author_approval ProjectLevelMrApprovals#merge_requests_author_approval}
	MergeRequestsAuthorApproval interface{} `field:"optional" json:"mergeRequestsAuthorApproval" yaml:"mergeRequestsAuthorApproval"`
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#merge_requests_disable_committers_approval ProjectLevelMrApprovals#merge_requests_disable_committers_approval}
	MergeRequestsDisableCommittersApproval interface{} `field:"optional" json:"mergeRequestsDisableCommittersApproval" yaml:"mergeRequestsDisableCommittersApproval"`
	// Set to `true` if you want to require authentication when approving a merge request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#require_password_to_approve ProjectLevelMrApprovals#require_password_to_approve}
	RequirePasswordToApprove interface{} `field:"optional" json:"requirePasswordToApprove" yaml:"requirePasswordToApprove"`
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	//
	// Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project_level_mr_approvals#reset_approvals_on_push ProjectLevelMrApprovals#reset_approvals_on_push}
	ResetApprovalsOnPush interface{} `field:"optional" json:"resetApprovalsOnPush" yaml:"resetApprovalsOnPush"`
}

