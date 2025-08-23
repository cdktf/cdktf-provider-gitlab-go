// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectVariableConfig struct {
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
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#key ProjectVariable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name or id of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#project ProjectVariable#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#value ProjectVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The description of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#description ProjectVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment scope of the variable.
	//
	// Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#environment_scope ProjectVariable#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// If set to `true`, the value of the variable will be hidden in the CI/CD User Interface.
	//
	// The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#hidden ProjectVariable#hidden}
	Hidden interface{} `field:"optional" json:"hidden" yaml:"hidden"`
	// If set to `true`, the value of the variable will be masked in job logs.
	//
	// The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#masked ProjectVariable#masked}
	Masked interface{} `field:"optional" json:"masked" yaml:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#protected ProjectVariable#protected}
	Protected interface{} `field:"optional" json:"protected" yaml:"protected"`
	// Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#raw ProjectVariable#raw}
	Raw interface{} `field:"optional" json:"raw" yaml:"raw"`
	// The type of a variable. Valid values are: `env_var`, `file`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_variable#variable_type ProjectVariable#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

