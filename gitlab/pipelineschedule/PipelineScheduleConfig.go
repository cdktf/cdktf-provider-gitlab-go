// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipelineschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineScheduleConfig struct {
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
	// The cron (e.g. `0 1 * * *`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#cron PipelineSchedule#cron}
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// The description of the pipeline schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#description PipelineSchedule#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name or id of the project to add the schedule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#project PipelineSchedule#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The branch/tag name to be triggered. This must be the full branch reference, for example: `refs/heads/main`, not `main`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#ref PipelineSchedule#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#active PipelineSchedule#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// The timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#cron_timezone PipelineSchedule#cron_timezone}
	CronTimezone *string `field:"optional" json:"cronTimezone" yaml:"cronTimezone"`
	// When set to `true`, the user represented by the token running Terraform will take ownership of the scheduled pipeline prior to editing it.
	//
	// This can help when managing scheduled pipeline drift when other users are making changes outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs/resources/pipeline_schedule#take_ownership PipelineSchedule#take_ownership}
	TakeOwnership interface{} `field:"optional" json:"takeOwnership" yaml:"takeOwnership"`
}

