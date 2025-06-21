// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagProtectionConfig struct {
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
	// The id of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/tag_protection#project TagProtection#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Name of the tag or wildcard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/tag_protection#tag TagProtection#tag}
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// allowed_to_create block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/tag_protection#allowed_to_create TagProtection#allowed_to_create}
	AllowedToCreate interface{} `field:"optional" json:"allowedToCreate" yaml:"allowedToCreate"`
	// Access levels allowed to create.
	//
	// Default value of `maintainer`. The default value is always sent if not provided in the configuration. Valid values are: `no one`, `developer`, `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/tag_protection#create_access_level TagProtection#create_access_level}
	CreateAccessLevel *string `field:"optional" json:"createAccessLevel" yaml:"createAccessLevel"`
}

