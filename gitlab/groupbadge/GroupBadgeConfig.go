// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupbadge

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupBadgeConfig struct {
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
	// The ID or URL-encoded path of the group to add the badge to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_badge#group GroupBadge#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The image url which will be presented on group overview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_badge#image_url GroupBadge#image_url}
	ImageUrl *string `field:"required" json:"imageUrl" yaml:"imageUrl"`
	// The url linked with the badge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_badge#link_url GroupBadge#link_url}
	LinkUrl *string `field:"required" json:"linkUrl" yaml:"linkUrl"`
	// The name of the badge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_badge#name GroupBadge#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

