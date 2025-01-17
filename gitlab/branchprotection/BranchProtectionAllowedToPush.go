// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection


type BranchProtectionAllowedToPush struct {
	// The ID of a GitLab deploy key allowed to perform the relevant action.
	//
	// Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/branch_protection#deploy_key_id BranchProtection#deploy_key_id}
	DeployKeyId *float64 `field:"optional" json:"deployKeyId" yaml:"deployKeyId"`
	// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/branch_protection#group_id BranchProtection#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/branch_protection#user_id BranchProtection#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

