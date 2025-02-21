// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectenvironments

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectEnvironmentsConfig struct {
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
	// The ID or full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_environments#project DataGitlabProjectEnvironments#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Return the environment with this name. Mutually exclusive with search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_environments#name DataGitlabProjectEnvironments#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Return list of environments matching the search criteria.
	//
	// Mutually exclusive with name. Must be at least 3 characters long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_environments#search DataGitlabProjectEnvironments#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// List all environments that match the specified state.
	//
	// Valid values are `available`, `stopping`, `stopped`. Returns all environments if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/data-sources/project_environments#states DataGitlabProjectEnvironments#states}
	States *string `field:"optional" json:"states" yaml:"states"`
}

