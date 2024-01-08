// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
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
	// The name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#name Group#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#path Group#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Default to Auto DevOps pipeline for all projects within this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#auto_devops_enabled Group#auto_devops_enabled}
	AutoDevopsEnabled interface{} `field:"optional" json:"autoDevopsEnabled" yaml:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#avatar Group#avatar}
	Avatar *string `field:"optional" json:"avatar" yaml:"avatar"`
	// The hash of the avatar image.
	//
	// Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#avatar_hash Group#avatar_hash}
	AvatarHash *string `field:"optional" json:"avatarHash" yaml:"avatarHash"`
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#default_branch_protection Group#default_branch_protection}
	DefaultBranchProtection *float64 `field:"optional" json:"defaultBranchProtection" yaml:"defaultBranchProtection"`
	// The group's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#description Group#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disable email notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#emails_disabled Group#emails_disabled}
	EmailsDisabled interface{} `field:"optional" json:"emailsDisabled" yaml:"emailsDisabled"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#extra_shared_runners_minutes_limit Group#extra_shared_runners_minutes_limit}
	ExtraSharedRunnersMinutesLimit *float64 `field:"optional" json:"extraSharedRunnersMinutesLimit" yaml:"extraSharedRunnersMinutesLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of IP addresses or subnet masks to restrict group access.
	//
	// Will be concatenated together into a comma separated string. Only allowed on top level groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#ip_restriction_ranges Group#ip_restriction_ranges}
	IpRestrictionRanges *[]*string `field:"optional" json:"ipRestrictionRanges" yaml:"ipRestrictionRanges"`
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#lfs_enabled Group#lfs_enabled}
	LfsEnabled interface{} `field:"optional" json:"lfsEnabled" yaml:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#membership_lock Group#membership_lock}
	MembershipLock interface{} `field:"optional" json:"membershipLock" yaml:"membershipLock"`
	// Disable the capability of a group from getting mentioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#mentions_disabled Group#mentions_disabled}
	MentionsDisabled interface{} `field:"optional" json:"mentionsDisabled" yaml:"mentionsDisabled"`
	// Id of the parent group (creates a nested group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#parent_id Group#parent_id}
	ParentId *float64 `field:"optional" json:"parentId" yaml:"parentId"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#prevent_forking_outside_group Group#prevent_forking_outside_group}
	PreventForkingOutsideGroup interface{} `field:"optional" json:"preventForkingOutsideGroup" yaml:"preventForkingOutsideGroup"`
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#project_creation_level Group#project_creation_level}
	ProjectCreationLevel *string `field:"optional" json:"projectCreationLevel" yaml:"projectCreationLevel"`
	// push_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#push_rules Group#push_rules}
	PushRules *GroupPushRules `field:"optional" json:"pushRules" yaml:"pushRules"`
	// Allow users to request member access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#request_access_enabled Group#request_access_enabled}
	RequestAccessEnabled interface{} `field:"optional" json:"requestAccessEnabled" yaml:"requestAccessEnabled"`
	// Require all users in this group to setup Two-factor authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#require_two_factor_authentication Group#require_two_factor_authentication}
	RequireTwoFactorAuthentication interface{} `field:"optional" json:"requireTwoFactorAuthentication" yaml:"requireTwoFactorAuthentication"`
	// Can be set by administrators only.
	//
	// Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#shared_runners_minutes_limit Group#shared_runners_minutes_limit}
	SharedRunnersMinutesLimit *float64 `field:"optional" json:"sharedRunnersMinutesLimit" yaml:"sharedRunnersMinutesLimit"`
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#shared_runners_setting Group#shared_runners_setting}
	SharedRunnersSetting *string `field:"optional" json:"sharedRunnersSetting" yaml:"sharedRunnersSetting"`
	// Prevent sharing a project with another group within this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#share_with_group_lock Group#share_with_group_lock}
	ShareWithGroupLock interface{} `field:"optional" json:"shareWithGroupLock" yaml:"shareWithGroupLock"`
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#subgroup_creation_level Group#subgroup_creation_level}
	SubgroupCreationLevel *string `field:"optional" json:"subgroupCreationLevel" yaml:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#two_factor_grace_period Group#two_factor_grace_period}
	TwoFactorGracePeriod *float64 `field:"optional" json:"twoFactorGracePeriod" yaml:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#visibility_level Group#visibility_level}
	VisibilityLevel *string `field:"optional" json:"visibilityLevel" yaml:"visibilityLevel"`
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.7.0/docs/resources/group#wiki_access_level Group#wiki_access_level}
	WikiAccessLevel *string `field:"optional" json:"wikiAccessLevel" yaml:"wikiAccessLevel"`
}

