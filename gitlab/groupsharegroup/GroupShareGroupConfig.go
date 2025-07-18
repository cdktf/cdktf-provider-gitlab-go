// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupsharegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupShareGroupConfig struct {
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
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_share_group#group_access GroupShareGroup#group_access}
	GroupAccess *string `field:"required" json:"groupAccess" yaml:"groupAccess"`
	// The id of the main group to be shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_share_group#group_id GroupShareGroup#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The id of the additional group with which the main group will be shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_share_group#share_group_id GroupShareGroup#share_group_id}
	ShareGroupId *float64 `field:"required" json:"shareGroupId" yaml:"shareGroupId"`
	// Share expiration date. Format: `YYYY-MM-DD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_share_group#expires_at GroupShareGroup#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The ID of a custom member role.
	//
	// Only available for Ultimate instances. If `member_role_id` is removed from the config, the group share will revert to a base role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_share_group#member_role_id GroupShareGroup#member_role_id}
	MemberRoleId *float64 `field:"optional" json:"memberRoleId" yaml:"memberRoleId"`
}

