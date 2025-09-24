// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabrepositoryfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabRepositoryFileConfig struct {
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
	// The full path of the file.
	//
	// It must be relative to the root of the project without a leading slash `/` or `./`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/repository_file#file_path DataGitlabRepositoryFile#file_path}
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The name or ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/repository_file#project DataGitlabRepositoryFile#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The name of branch, tag or commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/repository_file#ref DataGitlabRepositoryFile#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
}

