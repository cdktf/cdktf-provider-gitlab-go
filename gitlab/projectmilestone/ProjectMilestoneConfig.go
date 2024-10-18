// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectmilestone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectMilestoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#project ProjectMilestone#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The title of a milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#title ProjectMilestone#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The description of the milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#description ProjectMilestone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#due_date ProjectMilestone#due_date}
	DueDate *string `field:"optional" json:"dueDate" yaml:"dueDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#id ProjectMilestone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#start_date ProjectMilestone#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// The state of the milestone. Valid values are: `active`, `closed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_milestone#state ProjectMilestone#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

