// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupprotectedenvironment


type GroupProtectedEnvironmentApprovalRules struct {
	// Levels of access allowed to approve a deployment to this protected environment. Valid values are `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#access_level GroupProtectedEnvironment#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The ID of the group allowed to approve a deployment to this protected environment.
	//
	// TThe group must be a sub-group under the given group. This is mutually exclusive with user_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#group_id GroupProtectedEnvironment#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The number of approval required to allow deployment to this protected environment. This is mutually exclusive with user_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#required_approvals GroupProtectedEnvironment#required_approvals}
	RequiredApprovals *float64 `field:"optional" json:"requiredApprovals" yaml:"requiredApprovals"`
	// The ID of the user allowed to approve a deployment to this protected environment.
	//
	// The user must be a member of the group with Maintainer role or higher. This is mutually exclusive with group_id and required_approvals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#user_id GroupProtectedEnvironment#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

