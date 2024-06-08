// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojecttags

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectTagsConfig struct {
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
	// The ID or URL-encoded path of the project owned by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_tags#project DataGitlabProjectTags#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_tags#id DataGitlabProjectTags#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Return tags ordered by `name` or `updated` fields. Default is `updated`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_tags#order_by DataGitlabProjectTags#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Return list of tags matching the search criteria.
	//
	// You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_tags#search DataGitlabProjectTags#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Return tags sorted in `asc` or `desc` order. Default is `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_tags#sort DataGitlabProjectTags#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
}

