// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package complianceframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComplianceFrameworkConfig struct {
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
	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#color ComplianceFramework#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// Description for the compliance framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#description ComplianceFramework#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name for the compliance framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#name ComplianceFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Full path of the namespace to add the compliance framework to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#namespace_path ComplianceFramework#namespace_path}
	NamespacePath *string `field:"required" json:"namespacePath" yaml:"namespacePath"`
	// Set this compliance framework as the default framework for the group. Default: `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#default ComplianceFramework#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/compliance_framework#pipeline_configuration_full_path ComplianceFramework#pipeline_configuration_full_path}
	PipelineConfigurationFullPath *string `field:"optional" json:"pipelineConfigurationFullPath" yaml:"pipelineConfigurationFullPath"`
}

