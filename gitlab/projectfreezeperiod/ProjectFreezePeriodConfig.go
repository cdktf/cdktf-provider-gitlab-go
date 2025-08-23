// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	// End of the Freeze Period in cron format (for example, `0 2 * * *`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_freeze_period#freeze_end ProjectFreezePeriod#freeze_end}
	FreezeEnd *string `field:"required" json:"freezeEnd" yaml:"freezeEnd"`
	// Start of the Freeze Period in cron format (for example, `0 1 * * *`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_freeze_period#freeze_start ProjectFreezePeriod#freeze_start}
	FreezeStart *string `field:"required" json:"freezeStart" yaml:"freezeStart"`
	// The ID or path of the project to add the freeze period to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_freeze_period#project ProjectFreezePeriod#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_freeze_period#cron_timezone ProjectFreezePeriod#cron_timezone}
	CronTimezone *string `field:"optional" json:"cronTimezone" yaml:"cronTimezone"`
}

