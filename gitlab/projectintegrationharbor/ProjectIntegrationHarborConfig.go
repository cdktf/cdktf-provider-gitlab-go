// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationharbor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationHarborConfig struct {
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
	// Password for authentication with the Harbor server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#password ProjectIntegrationHarbor#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// ID of the GitLab project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#project ProjectIntegrationHarbor#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL-friendly Harbor project name. This project needs to already exist in Harbor. Example: `my_project_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#project_name ProjectIntegrationHarbor#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Harbor URL. Example: `http://harbor.example.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#url ProjectIntegrationHarbor#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Username for authentication with the Harbor server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#username ProjectIntegrationHarbor#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_harbor#use_inherited_settings ProjectIntegrationHarbor#use_inherited_settings}
	UseInheritedSettings interface{} `field:"optional" json:"useInheritedSettings" yaml:"useInheritedSettings"`
}

