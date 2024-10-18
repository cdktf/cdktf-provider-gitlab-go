// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectpushrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectPushRulesAConfig struct {
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
	// The ID or URL-encoded path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#project ProjectPushRulesA#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// All commit author emails must match this regex, e.g. `@my-company.com$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#author_email_regex ProjectPushRulesA#author_email_regex}
	AuthorEmailRegex *string `field:"optional" json:"authorEmailRegex" yaml:"authorEmailRegex"`
	// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#branch_name_regex ProjectPushRulesA#branch_name_regex}
	BranchNameRegex *string `field:"optional" json:"branchNameRegex" yaml:"branchNameRegex"`
	// Users can only push commits to this repository that were committed with one of their own verified emails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#commit_committer_check ProjectPushRulesA#commit_committer_check}
	CommitCommitterCheck interface{} `field:"optional" json:"commitCommitterCheck" yaml:"commitCommitterCheck"`
	// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#commit_committer_name_check ProjectPushRulesA#commit_committer_name_check}
	CommitCommitterNameCheck interface{} `field:"optional" json:"commitCommitterNameCheck" yaml:"commitCommitterNameCheck"`
	// No commit message is allowed to match this regex, e.g. `ssh\:\/\/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#commit_message_negative_regex ProjectPushRulesA#commit_message_negative_regex}
	CommitMessageNegativeRegex *string `field:"optional" json:"commitMessageNegativeRegex" yaml:"commitMessageNegativeRegex"`
	// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#commit_message_regex ProjectPushRulesA#commit_message_regex}
	CommitMessageRegex *string `field:"optional" json:"commitMessageRegex" yaml:"commitMessageRegex"`
	// Deny deleting a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#deny_delete_tag ProjectPushRulesA#deny_delete_tag}
	DenyDeleteTag interface{} `field:"optional" json:"denyDeleteTag" yaml:"denyDeleteTag"`
	// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#file_name_regex ProjectPushRulesA#file_name_regex}
	FileNameRegex *string `field:"optional" json:"fileNameRegex" yaml:"fileNameRegex"`
	// Maximum file size (MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#max_file_size ProjectPushRulesA#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Restrict commits by author (email) to existing GitLab users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#member_check ProjectPushRulesA#member_check}
	MemberCheck interface{} `field:"optional" json:"memberCheck" yaml:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#prevent_secrets ProjectPushRulesA#prevent_secrets}
	PreventSecrets interface{} `field:"optional" json:"preventSecrets" yaml:"preventSecrets"`
	// Reject commit when it’s not DCO certified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#reject_non_dco_commits ProjectPushRulesA#reject_non_dco_commits}
	RejectNonDcoCommits interface{} `field:"optional" json:"rejectNonDcoCommits" yaml:"rejectNonDcoCommits"`
	// Reject commit when it’s not signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/project_push_rules#reject_unsigned_commits ProjectPushRulesA#reject_unsigned_commits}
	RejectUnsignedCommits interface{} `field:"optional" json:"rejectUnsignedCommits" yaml:"rejectUnsignedCommits"`
}

