// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroupvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupVariableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/data-sources/group_variable#group DataGitlabGroupVariable#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/data-sources/group_variable#key DataGitlabGroupVariable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The environment scope of the variable.
	//
	// Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/data-sources/group_variable#environment_scope DataGitlabGroupVariable#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/data-sources/group_variable#id DataGitlabGroupVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

