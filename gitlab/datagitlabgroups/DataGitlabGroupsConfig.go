// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroups

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/groups#id DataGitlabGroups#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/groups#order_by DataGitlabGroups#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Search groups by name or path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/groups#search DataGitlabGroups#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Sort groups' list in asc or desc order. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/groups#sort DataGitlabGroups#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Limit to top level groups, excluding all subgroups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/data-sources/groups#top_level_only DataGitlabGroups#top_level_only}
	TopLevelOnly interface{} `field:"optional" json:"topLevelOnly" yaml:"topLevelOnly"`
}

