// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationredmine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationRedmineConfig struct {
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
	// The URL to the Redmine project issue to link to this GitLab project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/integration_redmine#issues_url IntegrationRedmine#issues_url}
	IssuesUrl *string `field:"required" json:"issuesUrl" yaml:"issuesUrl"`
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/integration_redmine#new_issue_url IntegrationRedmine#new_issue_url}
	NewIssueUrl *string `field:"required" json:"newIssueUrl" yaml:"newIssueUrl"`
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/integration_redmine#project IntegrationRedmine#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL to the Redmine project to link to this GitLab project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/integration_redmine#project_url IntegrationRedmine#project_url}
	ProjectUrl *string `field:"required" json:"projectUrl" yaml:"projectUrl"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/integration_redmine#use_inherited_settings IntegrationRedmine#use_inherited_settings}
	UseInheritedSettings interface{} `field:"optional" json:"useInheritedSettings" yaml:"useInheritedSettings"`
}

