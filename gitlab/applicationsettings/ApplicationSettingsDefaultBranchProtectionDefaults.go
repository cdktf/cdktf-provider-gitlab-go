// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings


type ApplicationSettingsDefaultBranchProtectionDefaults struct {
	// An array of access levels allowed to merge. Supports Developer (30) or Maintainer (40).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/application_settings#allowed_to_merge ApplicationSettings#allowed_to_merge}
	AllowedToMerge *[]*float64 `field:"optional" json:"allowedToMerge" yaml:"allowedToMerge"`
	// An array of access levels allowed to push. Supports Developer (30) or Maintainer (40).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/application_settings#allowed_to_push ApplicationSettings#allowed_to_push}
	AllowedToPush *[]*float64 `field:"optional" json:"allowedToPush" yaml:"allowedToPush"`
	// Allow force push for all users with push access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/application_settings#allow_force_push ApplicationSettings#allow_force_push}
	AllowForcePush interface{} `field:"optional" json:"allowForcePush" yaml:"allowForcePush"`
	// Allow developers to initial push.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/application_settings#developer_can_initial_push ApplicationSettings#developer_can_initial_push}
	DeveloperCanInitialPush interface{} `field:"optional" json:"developerCanInitialPush" yaml:"developerCanInitialPush"`
}

