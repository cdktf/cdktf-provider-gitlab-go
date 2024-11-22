// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploykeyenable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeployKeyEnableConfig struct {
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
	// The Gitlab key id for the pre-existing deploy key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#key_id DeployKeyEnable#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The name or id of the project to add the deploy key to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#project DeployKeyEnable#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Can deploy key push to the project's repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#can_push DeployKeyEnable#can_push}
	CanPush interface{} `field:"optional" json:"canPush" yaml:"canPush"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#id DeployKeyEnable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Deploy key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#key DeployKeyEnable#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Deploy key's title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/deploy_key_enable#title DeployKeyEnable#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

