// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectMembershipConfig struct {
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
	// The full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#full_path DataGitlabProjectMembership#full_path}
	FullPath *string `field:"optional" json:"fullPath" yaml:"fullPath"`
	// Return all project members including members through ancestor groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#inherited DataGitlabProjectMembership#inherited}
	Inherited interface{} `field:"optional" json:"inherited" yaml:"inherited"`
	// The ID or full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#project DataGitlabProjectMembership#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#project_id DataGitlabProjectMembership#project_id}
	ProjectId *float64 `field:"optional" json:"projectId" yaml:"projectId"`
	// A query string to search for members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#query DataGitlabProjectMembership#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// List of user ids to filter members by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_membership#user_ids DataGitlabProjectMembership#user_ids}
	UserIds *[]*float64 `field:"optional" json:"userIds" yaml:"userIds"`
}

