// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationcustomissuetracker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationCustomIssueTrackerConfig struct {
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
	// The URL to view an issue in the external issue tracker. Must contain :id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_custom_issue_tracker#issues_url ProjectIntegrationCustomIssueTracker#issues_url}
	IssuesUrl *string `field:"required" json:"issuesUrl" yaml:"issuesUrl"`
	// The ID or full path of the project for the custom issue tracker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_custom_issue_tracker#project ProjectIntegrationCustomIssueTracker#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL to the project in the external issue tracker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_custom_issue_tracker#project_url ProjectIntegrationCustomIssueTracker#project_url}
	ProjectUrl *string `field:"required" json:"projectUrl" yaml:"projectUrl"`
}

