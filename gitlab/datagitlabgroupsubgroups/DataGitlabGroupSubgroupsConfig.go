// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroupsubgroups

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupSubgroupsConfig struct {
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
	// The ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#group_id DataGitlabGroupSubgroups#group_id}
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// Show all the groups you have access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#all_available DataGitlabGroupSubgroups#all_available}
	AllAvailable interface{} `field:"optional" json:"allAvailable" yaml:"allAvailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#id DataGitlabGroupSubgroups#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Limit to groups where current user has at least this access level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#min_access_level DataGitlabGroupSubgroups#min_access_level}
	MinAccessLevel *string `field:"optional" json:"minAccessLevel" yaml:"minAccessLevel"`
	// Order groups by name, path or id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#order_by DataGitlabGroupSubgroups#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Limit to groups explicitly owned by the current user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#owned DataGitlabGroupSubgroups#owned}
	Owned interface{} `field:"optional" json:"owned" yaml:"owned"`
	// Return the list of authorized groups matching the search criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#search DataGitlabGroupSubgroups#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Skip the group IDs passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#skip_groups DataGitlabGroupSubgroups#skip_groups}
	SkipGroups *[]*float64 `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Order groups in asc or desc order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#sort DataGitlabGroupSubgroups#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Include group statistics (administrators only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#statistics DataGitlabGroupSubgroups#statistics}
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
	// Include custom attributes in response (administrators only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/data-sources/group_subgroups#with_custom_attributes DataGitlabGroupSubgroups#with_custom_attributes}
	WithCustomAttributes interface{} `field:"optional" json:"withCustomAttributes" yaml:"withCustomAttributes"`
}

