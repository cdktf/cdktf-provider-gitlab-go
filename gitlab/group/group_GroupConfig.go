package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#name Group#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#path Group#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#auto_devops_enabled Group#auto_devops_enabled}
	AutoDevopsEnabled interface{} `field:"optional" json:"autoDevopsEnabled" yaml:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#avatar Group#avatar}
	Avatar *string `field:"optional" json:"avatar" yaml:"avatar"`
	// The hash of the avatar image.
	//
	// Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#avatar_hash Group#avatar_hash}
	AvatarHash *string `field:"optional" json:"avatarHash" yaml:"avatarHash"`
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#default_branch_protection Group#default_branch_protection}
	DefaultBranchProtection *float64 `field:"optional" json:"defaultBranchProtection" yaml:"defaultBranchProtection"`
	// The description of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#description Group#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defaults to false. Disable email notifications.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#emails_disabled Group#emails_disabled}
	EmailsDisabled interface{} `field:"optional" json:"emailsDisabled" yaml:"emailsDisabled"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#extra_shared_runners_minutes_limit Group#extra_shared_runners_minutes_limit}
	ExtraSharedRunnersMinutesLimit *float64 `field:"optional" json:"extraSharedRunnersMinutesLimit" yaml:"extraSharedRunnersMinutesLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of IP addresses or subnet masks to restrict group access.
	//
	// Will be concatenated together into a comma separated string. Only allowed on top level groups.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#ip_restriction_ranges Group#ip_restriction_ranges}
	IpRestrictionRanges *[]*string `field:"optional" json:"ipRestrictionRanges" yaml:"ipRestrictionRanges"`
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#lfs_enabled Group#lfs_enabled}
	LfsEnabled interface{} `field:"optional" json:"lfsEnabled" yaml:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#membership_lock Group#membership_lock}
	MembershipLock interface{} `field:"optional" json:"membershipLock" yaml:"membershipLock"`
	// Defaults to false. Disable the capability of a group from getting mentioned.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#mentions_disabled Group#mentions_disabled}
	MentionsDisabled interface{} `field:"optional" json:"mentionsDisabled" yaml:"mentionsDisabled"`
	// Id of the parent group (creates a nested group).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#parent_id Group#parent_id}
	ParentId *float64 `field:"optional" json:"parentId" yaml:"parentId"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#prevent_forking_outside_group Group#prevent_forking_outside_group}
	PreventForkingOutsideGroup interface{} `field:"optional" json:"preventForkingOutsideGroup" yaml:"preventForkingOutsideGroup"`
	// Defaults to maintainer. Determine if developers can create projects in the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#project_creation_level Group#project_creation_level}
	ProjectCreationLevel *string `field:"optional" json:"projectCreationLevel" yaml:"projectCreationLevel"`
	// Defaults to false. Allow users to request member access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#request_access_enabled Group#request_access_enabled}
	RequestAccessEnabled interface{} `field:"optional" json:"requestAccessEnabled" yaml:"requestAccessEnabled"`
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#require_two_factor_authentication Group#require_two_factor_authentication}
	RequireTwoFactorAuthentication interface{} `field:"optional" json:"requireTwoFactorAuthentication" yaml:"requireTwoFactorAuthentication"`
	// Can be set by administrators only.
	//
	// Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#shared_runners_minutes_limit Group#shared_runners_minutes_limit}
	SharedRunnersMinutesLimit *float64 `field:"optional" json:"sharedRunnersMinutesLimit" yaml:"sharedRunnersMinutesLimit"`
	// Defaults to false. Prevent sharing a project with another group within this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#share_with_group_lock Group#share_with_group_lock}
	ShareWithGroupLock interface{} `field:"optional" json:"shareWithGroupLock" yaml:"shareWithGroupLock"`
	// Defaults to owner. Allowed to create subgroups.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#subgroup_creation_level Group#subgroup_creation_level}
	SubgroupCreationLevel *string `field:"optional" json:"subgroupCreationLevel" yaml:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#two_factor_grace_period Group#two_factor_grace_period}
	TwoFactorGracePeriod *float64 `field:"optional" json:"twoFactorGracePeriod" yaml:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group#visibility_level Group#visibility_level}
	VisibilityLevel *string `field:"optional" json:"visibilityLevel" yaml:"visibilityLevel"`
}

