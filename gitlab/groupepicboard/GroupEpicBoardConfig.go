// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupepicboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupEpicBoardConfig struct {
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
	// The ID or URL-encoded path of the group owned by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/group_epic_board#group GroupEpicBoard#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the board.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/group_epic_board#name GroupEpicBoard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// lists block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/group_epic_board#lists GroupEpicBoard#lists}
	Lists interface{} `field:"optional" json:"lists" yaml:"lists"`
}

