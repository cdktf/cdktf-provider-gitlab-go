// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryFileConfig struct {
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
	// Name of the branch to which to commit to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#branch RepositoryFile#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// File content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#content RepositoryFile#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file content encoding. Valid values are: `base64`, `text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#encoding RepositoryFile#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// The full path of the file.
	//
	// It must be relative to the root of the project without a leading slash `/` or `./`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#file_path RepositoryFile#file_path}
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The name or ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#project RepositoryFile#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Email of the commit author.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#author_email RepositoryFile#author_email}
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Name of the commit author.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#author_name RepositoryFile#author_name}
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#commit_message RepositoryFile#commit_message}
	CommitMessage *string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// Create commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#create_commit_message RepositoryFile#create_commit_message}
	CreateCommitMessage *string `field:"optional" json:"createCommitMessage" yaml:"createCommitMessage"`
	// Delete Commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#delete_commit_message RepositoryFile#delete_commit_message}
	DeleteCommitMessage *string `field:"optional" json:"deleteCommitMessage" yaml:"deleteCommitMessage"`
	// Enables or disables the execute flag on the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#execute_filemode RepositoryFile#execute_filemode}
	ExecuteFilemode interface{} `field:"optional" json:"executeFilemode" yaml:"executeFilemode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#id RepositoryFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable overwriting existing files, defaults to `false`.
	//
	// This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#overwrite_on_create RepositoryFile#overwrite_on_create}
	OverwriteOnCreate interface{} `field:"optional" json:"overwriteOnCreate" yaml:"overwriteOnCreate"`
	// Name of the branch to start the new commit from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#start_branch RepositoryFile#start_branch}
	StartBranch *string `field:"optional" json:"startBranch" yaml:"startBranch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#timeouts RepositoryFile#timeouts}
	Timeouts *RepositoryFileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Update commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file#update_commit_message RepositoryFile#update_commit_message}
	UpdateCommitMessage *string `field:"optional" json:"updateCommitMessage" yaml:"updateCommitMessage"`
}

