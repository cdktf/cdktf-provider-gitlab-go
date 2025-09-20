// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouplevelmrapprovals

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupLevelMrApprovalsConfig struct {
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
	// The ID or URL-encoded path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#group GroupLevelMrApprovals#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Allow or prevent authors from self approving merge requests; `true` means authors can self approve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#allow_author_approval GroupLevelMrApprovals#allow_author_approval}
	AllowAuthorApproval interface{} `field:"optional" json:"allowAuthorApproval" yaml:"allowAuthorApproval"`
	// Allow or prevent committers from self approving merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#allow_committer_approval GroupLevelMrApprovals#allow_committer_approval}
	AllowCommitterApproval interface{} `field:"optional" json:"allowCommitterApproval" yaml:"allowCommitterApproval"`
	// Allow or prevent overriding approvers per merge request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#allow_overrides_to_approver_list_per_merge_request GroupLevelMrApprovals#allow_overrides_to_approver_list_per_merge_request}
	AllowOverridesToApproverListPerMergeRequest interface{} `field:"optional" json:"allowOverridesToApproverListPerMergeRequest" yaml:"allowOverridesToApproverListPerMergeRequest"`
	// Set to true if the group merge request approval settings should not be reset to their pre-terraform defaults on destroy.
	//
	// You will need to apply the resource with the new setting before destroying the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#keep_settings_on_destroy GroupLevelMrApprovals#keep_settings_on_destroy}
	KeepSettingsOnDestroy interface{} `field:"optional" json:"keepSettingsOnDestroy" yaml:"keepSettingsOnDestroy"`
	// Require approver to authenticate before adding the approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#require_reauthentication_to_approve GroupLevelMrApprovals#require_reauthentication_to_approve}
	RequireReauthenticationToApprove interface{} `field:"optional" json:"requireReauthenticationToApprove" yaml:"requireReauthenticationToApprove"`
	// Retain approval count on a new push.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals#retain_approvals_on_push GroupLevelMrApprovals#retain_approvals_on_push}
	RetainApprovalsOnPush interface{} `field:"optional" json:"retainApprovalsOnPush" yaml:"retainApprovalsOnPush"`
}

