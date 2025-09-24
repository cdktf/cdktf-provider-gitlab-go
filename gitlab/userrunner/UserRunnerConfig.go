// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userrunner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserRunnerConfig struct {
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
	// The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#runner_type UserRunner#runner_type}
	RunnerType *string `field:"required" json:"runnerType" yaml:"runnerType"`
	// The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#access_level UserRunner#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// Description of the runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#description UserRunner#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the group that the runner is created in. Required if runner_type is group_type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#group_id UserRunner#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Specifies if the runner should be locked for the current project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#locked UserRunner#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// Free-form maintenance notes for the runner (1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#maintenance_note UserRunner#maintenance_note}
	MaintenanceNote *string `field:"optional" json:"maintenanceNote" yaml:"maintenanceNote"`
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs.
	//
	// Must be at least 600 (10 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#maximum_timeout UserRunner#maximum_timeout}
	MaximumTimeout *float64 `field:"optional" json:"maximumTimeout" yaml:"maximumTimeout"`
	// Specifies if the runner should ignore new jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#paused UserRunner#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The ID of the project that the runner is created in. Required if runner_type is project_type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#project_id UserRunner#project_id}
	ProjectId *float64 `field:"optional" json:"projectId" yaml:"projectId"`
	// A list of runner tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#tag_list UserRunner#tag_list}
	TagList *[]*string `field:"optional" json:"tagList" yaml:"tagList"`
	// Specifies if the runner should handle untagged jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user_runner#untagged UserRunner#untagged}
	Untagged interface{} `field:"optional" json:"untagged" yaml:"untagged"`
}

