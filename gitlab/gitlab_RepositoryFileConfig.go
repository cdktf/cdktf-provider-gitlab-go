// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryFileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#branch RepositoryFile#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// Commit message.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#commit_message RepositoryFile#commit_message}
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// File content.
	//
	// If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#content RepositoryFile#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The full path of the file.
	//
	// It must be relative to the root of the project without a leading slash `/`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#file_path RepositoryFile#file_path}
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The name or ID of the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#project RepositoryFile#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Email of the commit author.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#author_email RepositoryFile#author_email}
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Name of the commit author.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#author_name RepositoryFile#author_name}
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#execute_filemode RepositoryFile#execute_filemode}
	ExecuteFilemode interface{} `field:"optional" json:"executeFilemode" yaml:"executeFilemode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#id RepositoryFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the branch to start the new commit from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#start_branch RepositoryFile#start_branch}
	StartBranch *string `field:"optional" json:"startBranch" yaml:"startBranch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/repository_file#timeouts RepositoryFile#timeouts}
	Timeouts *RepositoryFileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

