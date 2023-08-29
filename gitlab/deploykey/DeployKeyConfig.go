// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploykey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeployKeyConfig struct {
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
	// The public ssh key body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/deploy_key#key DeployKey#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name or id of the project to add the deploy key to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/deploy_key#project DeployKey#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// A title to describe the deploy key with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/deploy_key#title DeployKey#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Allow this deploy key to be used to push changes to the project. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/deploy_key#can_push DeployKey#can_push}
	CanPush interface{} `field:"optional" json:"canPush" yaml:"canPush"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/deploy_key#id DeployKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

