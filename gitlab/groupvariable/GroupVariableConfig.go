// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupVariableConfig struct {
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
	// The name or id of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#group GroupVariable#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#key GroupVariable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#value GroupVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The description of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#description GroupVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment scope of the variable.
	//
	// Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#environment_scope GroupVariable#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// If set to `true`, the value of the variable will be hidden in the CI/CD User Interface.
	//
	// The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#hidden GroupVariable#hidden}
	Hidden interface{} `field:"optional" json:"hidden" yaml:"hidden"`
	// If set to `true`, the value of the variable will be masked in job logs.
	//
	// The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#mask-a-cicd-variable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#masked GroupVariable#masked}
	Masked interface{} `field:"optional" json:"masked" yaml:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#protected GroupVariable#protected}
	Protected interface{} `field:"optional" json:"protected" yaml:"protected"`
	// Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#raw GroupVariable#raw}
	Raw interface{} `field:"optional" json:"raw" yaml:"raw"`
	// The type of a variable. Valid values are: `env_var`, `file`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_variable#variable_type GroupVariable#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

