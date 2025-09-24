// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationpipelinesemail

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationPipelinesEmailConfig struct {
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
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_pipelines_email#project ProjectIntegrationPipelinesEmail#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// ) email addresses where notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_pipelines_email#recipients ProjectIntegrationPipelinesEmail#recipients}
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_pipelines_email#branches_to_be_notified ProjectIntegrationPipelinesEmail#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_pipelines_email#id ProjectIntegrationPipelinesEmail#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Notify only broken pipelines. Default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_pipelines_email#notify_only_broken_pipelines ProjectIntegrationPipelinesEmail#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
}

