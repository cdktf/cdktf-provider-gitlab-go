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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/resources/deploy_key_enable#key_id DeployKeyEnable#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The name or id of the project to add the deploy key to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/resources/deploy_key_enable#project DeployKeyEnable#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Can deploy key push to the project's repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/resources/deploy_key_enable#can_push DeployKeyEnable#can_push}
	CanPush interface{} `field:"optional" json:"canPush" yaml:"canPush"`
}

