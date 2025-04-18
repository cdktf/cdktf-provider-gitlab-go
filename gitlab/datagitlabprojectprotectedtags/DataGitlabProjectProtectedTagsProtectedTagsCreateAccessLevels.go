// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedtags


type DataGitlabProjectProtectedTagsProtectedTagsCreateAccessLevels struct {
	// The ID of a GitLab group allowed to perform the relevant action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/data-sources/project_protected_tags#group_id DataGitlabProjectProtectedTags#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/data-sources/project_protected_tags#user_id DataGitlabProjectProtectedTags#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

