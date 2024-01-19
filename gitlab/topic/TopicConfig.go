// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package topic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TopicConfig struct {
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
	// The topic's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#name Topic#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#avatar Topic#avatar}
	Avatar *string `field:"optional" json:"avatar" yaml:"avatar"`
	// The hash of the avatar image.
	//
	// Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#avatar_hash Topic#avatar_hash}
	AvatarHash *string `field:"optional" json:"avatarHash" yaml:"avatarHash"`
	// A text describing the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#description Topic#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#id Topic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Empty the topics fields instead of deleting it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#soft_destroy Topic#soft_destroy}
	SoftDestroy interface{} `field:"optional" json:"softDestroy" yaml:"softDestroy"`
	// The topic's description. Requires at least GitLab 15.0 for which it's a required argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/topic#title Topic#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

