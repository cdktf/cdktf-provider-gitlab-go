// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package label

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LabelConfig struct {
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
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/label#color Label#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// The name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/label#name Label#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or id of the project to add the label to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/label#project Label#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The description of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/label#description Label#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

