// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabrepositorytree

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabRepositoryTreeConfig struct {
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
	// The ID or full path of the project owned by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/repository_tree#project DataGitlabRepositoryTree#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The name of a repository branch or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/repository_tree#ref DataGitlabRepositoryTree#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// The path inside repository. Used to get content of subdirectories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/repository_tree#path DataGitlabRepositoryTree#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Boolean value used to get a recursive tree (false by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/repository_tree#recursive DataGitlabRepositoryTree#recursive}
	Recursive interface{} `field:"optional" json:"recursive" yaml:"recursive"`
}

