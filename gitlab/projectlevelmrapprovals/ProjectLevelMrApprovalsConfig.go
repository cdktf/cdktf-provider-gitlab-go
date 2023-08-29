// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#project ProjectLevelMrApprovals#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Set to `true` to disable overriding approvers per merge request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#disable_overriding_approvers_per_merge_request ProjectLevelMrApprovals#disable_overriding_approvers_per_merge_request}
	DisableOverridingApproversPerMergeRequest interface{} `field:"optional" json:"disableOverridingApproversPerMergeRequest" yaml:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#merge_requests_author_approval ProjectLevelMrApprovals#merge_requests_author_approval}
	MergeRequestsAuthorApproval interface{} `field:"optional" json:"mergeRequestsAuthorApproval" yaml:"mergeRequestsAuthorApproval"`
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#merge_requests_disable_committers_approval ProjectLevelMrApprovals#merge_requests_disable_committers_approval}
	MergeRequestsDisableCommittersApproval interface{} `field:"optional" json:"mergeRequestsDisableCommittersApproval" yaml:"mergeRequestsDisableCommittersApproval"`
	// Set to `true` to require authentication to approve merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#require_password_to_approve ProjectLevelMrApprovals#require_password_to_approve}
	RequirePasswordToApprove interface{} `field:"optional" json:"requirePasswordToApprove" yaml:"requirePasswordToApprove"`
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch.
	//
	// Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#reset_approvals_on_push ProjectLevelMrApprovals#reset_approvals_on_push}
	ResetApprovalsOnPush interface{} `field:"optional" json:"resetApprovalsOnPush" yaml:"resetApprovalsOnPush"`
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset_approvals_on_push is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/project_level_mr_approvals#selective_code_owner_removals ProjectLevelMrApprovals#selective_code_owner_removals}
	SelectiveCodeOwnerRemovals interface{} `field:"optional" json:"selectiveCodeOwnerRemovals" yaml:"selectiveCodeOwnerRemovals"`
}

