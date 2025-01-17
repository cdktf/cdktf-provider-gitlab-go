// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectprotectedenvironment


type ProjectProtectedEnvironmentApprovalRules struct {
	// Levels of access allowed to approve a deployment to this protected environment.
	//
	// Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_protected_environment#access_level ProjectProtectedEnvironment#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The ID of the group allowed to approve a deployment to this protected environment.
	//
	// The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_protected_environment#group_id ProjectProtectedEnvironment#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Group inheritance allows deploy access levels to take inherited group membership into account.
	//
	// Valid values are `0`, `1`. `0` => Direct group membership only, `1` => All inherited groups. Default: `0`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_protected_environment#group_inheritance_type ProjectProtectedEnvironment#group_inheritance_type}
	GroupInheritanceType *float64 `field:"optional" json:"groupInheritanceType" yaml:"groupInheritanceType"`
	// The number of approval required to allow deployment to this protected environment. This is mutually exclusive with user_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_protected_environment#required_approvals ProjectProtectedEnvironment#required_approvals}
	RequiredApprovals *float64 `field:"optional" json:"requiredApprovals" yaml:"requiredApprovals"`
	// The ID of the user allowed to approve a deployment to this protected environment.
	//
	// The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_protected_environment#user_id ProjectProtectedEnvironment#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

