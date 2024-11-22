// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedbranch


type DataGitlabProjectProtectedBranchPushAccessLevels struct {
	// The ID of a GitLab deploy key allowed to perform the relevant action.
	//
	// Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/data-sources/project_protected_branch#deploy_key_id DataGitlabProjectProtectedBranch#deploy_key_id}
	DeployKeyId *float64 `field:"optional" json:"deployKeyId" yaml:"deployKeyId"`
	// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/data-sources/project_protected_branch#group_id DataGitlabProjectProtectedBranch#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/data-sources/project_protected_branch#user_id DataGitlabProjectProtectedBranch#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

