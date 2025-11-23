// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchConfig struct {
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
	// The name for this branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.1/docs/resources/branch#name Branch#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID or full path of the project which the branch is created against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.1/docs/resources/branch#project Branch#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The ref which the branch is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.1/docs/resources/branch#ref Branch#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// Indicates whether the branch is kept once the resource destroyed (must be applied before a destroy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.1/docs/resources/branch#keep_on_destroy Branch#keep_on_destroy}
	KeepOnDestroy interface{} `field:"optional" json:"keepOnDestroy" yaml:"keepOnDestroy"`
}

