// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagprotection


type TagProtectionAllowedToCreate struct {
	// Access levels allowed to create protected tags. Valid values are: `no one`, `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/tag_protection#access_level TagProtection#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/tag_protection#group_id TagProtection#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/tag_protection#user_id TagProtection#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

