// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedbranches


type DataGitlabProjectProtectedBranchesProtectedBranchesPushAccessLevels struct {
	// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/project_protected_branches#group_id DataGitlabProjectProtectedBranches#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/project_protected_branches#user_id DataGitlabProjectProtectedBranches#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

