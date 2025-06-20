// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancevariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstanceVariableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#key InstanceVariable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#value InstanceVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The description of the variable. Maximum of 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#description InstanceVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#id InstanceVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set to `true`, the value of the variable will be hidden in job logs.
	//
	// The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#masked InstanceVariable#masked}
	Masked interface{} `field:"optional" json:"masked" yaml:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
	//
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#protected InstanceVariable#protected}
	Protected interface{} `field:"optional" json:"protected" yaml:"protected"`
	// Whether the variable is treated as a raw string.
	//
	// Default: false. When true, variables in the value are not expanded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#raw InstanceVariable#raw}
	Raw interface{} `field:"optional" json:"raw" yaml:"raw"`
	// The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/instance_variable#variable_type InstanceVariable#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

