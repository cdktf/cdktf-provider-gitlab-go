// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package release

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReleaseConfig struct {
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
	// The ID or full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#project Release#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The tag where the release is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#tag_name Release#tag_name}
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
	// The release assets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#assets Release#assets}
	Assets *ReleaseAssets `field:"optional" json:"assets" yaml:"assets"`
	// The description of the release. You can use Markdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#description Release#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#milestones Release#milestones}
	Milestones *[]*string `field:"optional" json:"milestones" yaml:"milestones"`
	// The name of the release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#name Release#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If a tag specified in tag_name doesn't exist, the release is created from ref and tagged with tag_name.
	//
	// It can be a commit SHA, another tag name, or a branch name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#ref Release#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Date and time for the release.
	//
	// Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#released_at Release#released_at}
	ReleasedAt *string `field:"optional" json:"releasedAt" yaml:"releasedAt"`
	// Message to use if creating a new annotated tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/release#tag_message Release#tag_message}
	TagMessage *string `field:"optional" json:"tagMessage" yaml:"tagMessage"`
}

