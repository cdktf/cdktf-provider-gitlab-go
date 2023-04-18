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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#cron PipelineSchedule#cron}
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// The description of the pipeline schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#description PipelineSchedule#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name or id of the project to add the schedule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#project PipelineSchedule#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The branch/tag name to be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#ref PipelineSchedule#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#active PipelineSchedule#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// The timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#cron_timezone PipelineSchedule#cron_timezone}
	CronTimezone *string `field:"optional" json:"cronTimezone" yaml:"cronTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/pipeline_schedule#id PipelineSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

