package projectfreezeperiod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectFreezePeriodConfig struct {
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
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_freeze_period#freeze_end ProjectFreezePeriod#freeze_end}
	FreezeEnd *string `field:"required" json:"freezeEnd" yaml:"freezeEnd"`
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_freeze_period#freeze_start ProjectFreezePeriod#freeze_start}
	FreezeStart *string `field:"required" json:"freezeStart" yaml:"freezeStart"`
	// The id of the project to add the schedule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_freeze_period#project_id ProjectFreezePeriod#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_freeze_period#cron_timezone ProjectFreezePeriod#cron_timezone}
	CronTimezone *string `field:"optional" json:"cronTimezone" yaml:"cronTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_freeze_period#id ProjectFreezePeriod#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

