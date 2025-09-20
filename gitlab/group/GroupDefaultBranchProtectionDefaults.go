// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group


type GroupDefaultBranchProtectionDefaults struct {
	// An array of access levels allowed to merge. Valid values are: `developer`, `maintainer`, `no one`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group#allowed_to_merge Group#allowed_to_merge}
	AllowedToMerge *[]*string `field:"optional" json:"allowedToMerge" yaml:"allowedToMerge"`
	// An array of access levels allowed to push. Valid values are: `developer`, `maintainer`, `no one`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group#allowed_to_push Group#allowed_to_push}
	AllowedToPush *[]*string `field:"optional" json:"allowedToPush" yaml:"allowedToPush"`
	// Allow force push for all users with push access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group#allow_force_push Group#allow_force_push}
	AllowForcePush interface{} `field:"optional" json:"allowForcePush" yaml:"allowForcePush"`
	// Allow developers to initial push.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group#developer_can_initial_push Group#developer_can_initial_push}
	DeveloperCanInitialPush interface{} `field:"optional" json:"developerCanInitialPush" yaml:"developerCanInitialPush"`
}

