// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMembershipConfig struct {
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
	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#access_level GroupMembership#access_level}
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// The ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#group_id GroupMembership#group_id}
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#user_id GroupMembership#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#expires_at GroupMembership#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The ID of a custom member role. Only available for Ultimate instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#member_role_id GroupMembership#member_role_id}
	MemberRoleId *float64 `field:"optional" json:"memberRoleId" yaml:"memberRoleId"`
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped.
	//
	// Only used during a destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#skip_subresources_on_destroy GroupMembership#skip_subresources_on_destroy}
	SkipSubresourcesOnDestroy interface{} `field:"optional" json:"skipSubresourcesOnDestroy" yaml:"skipSubresourcesOnDestroy"`
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project.
	//
	// Only used during a destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_membership#unassign_issuables_on_destroy GroupMembership#unassign_issuables_on_destroy}
	UnassignIssuablesOnDestroy interface{} `field:"optional" json:"unassignIssuablesOnDestroy" yaml:"unassignIssuablesOnDestroy"`
}

