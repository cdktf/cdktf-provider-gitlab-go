// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectprotectedenvironment


type ProjectProtectedEnvironmentDeployAccessLevels struct {
	// Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.4.0/docs/resources/project_protected_environment#access_level ProjectProtectedEnvironment#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The ID of the group allowed to deploy to this protected environment.
	//
	// The project must be shared with the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.4.0/docs/resources/project_protected_environment#group_id ProjectProtectedEnvironment#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of the user allowed to deploy to this protected environment.
	//
	// The user must be a member of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.4.0/docs/resources/project_protected_environment#user_id ProjectProtectedEnvironment#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

