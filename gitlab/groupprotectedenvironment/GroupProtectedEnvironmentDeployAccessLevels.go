// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupprotectedenvironment


type GroupProtectedEnvironmentDeployAccessLevels struct {
	// Levels of access required to deploy to this protected environment.
	//
	// Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/group_protected_environment#access_level GroupProtectedEnvironment#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The ID of the group allowed to deploy to this protected environment.
	//
	// The group must be a sub-group under the given group. Mutually exclusive with `access_level` and `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/group_protected_environment#group_id GroupProtectedEnvironment#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Group inheritance allows deploy access levels to take inherited group membership into account.
	//
	// Valid values are `0`, `1`. `0` => Direct group membership only, `1` => All inherited groups. Default: `0`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/group_protected_environment#group_inheritance_type GroupProtectedEnvironment#group_inheritance_type}
	GroupInheritanceType *float64 `field:"optional" json:"groupInheritanceType" yaml:"groupInheritanceType"`
	// The ID of the user allowed to deploy to this protected environment.
	//
	// The user must be a member of the group with Maintainer role or higher. Mutually exclusive with `access_level` and `group_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/group_protected_environment#user_id GroupProtectedEnvironment#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

