// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedtag


type DataGitlabProjectProtectedTagCreateAccessLevels struct {
	// The ID of a GitLab group allowed to perform the relevant action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_protected_tag#group_id DataGitlabProjectProtectedTag#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_protected_tag#user_id DataGitlabProjectProtectedTag#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

