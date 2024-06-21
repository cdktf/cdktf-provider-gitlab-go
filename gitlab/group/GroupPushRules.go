// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group


type GroupPushRules struct {
	// All commit author emails must match this regex, e.g. `@my-company.com$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#author_email_regex Group#author_email_regex}
	AuthorEmailRegex *string `field:"optional" json:"authorEmailRegex" yaml:"authorEmailRegex"`
	// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#branch_name_regex Group#branch_name_regex}
	BranchNameRegex *string `field:"optional" json:"branchNameRegex" yaml:"branchNameRegex"`
	// Only commits pushed using verified emails are allowed.
	//
	// **Note** This attribute is only supported in GitLab versions >= 16.4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#commit_committer_check Group#commit_committer_check}
	CommitCommitterCheck interface{} `field:"optional" json:"commitCommitterCheck" yaml:"commitCommitterCheck"`
	// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#commit_committer_name_check Group#commit_committer_name_check}
	CommitCommitterNameCheck interface{} `field:"optional" json:"commitCommitterNameCheck" yaml:"commitCommitterNameCheck"`
	// No commit message is allowed to match this regex, for example `ssh\:\/\/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#commit_message_negative_regex Group#commit_message_negative_regex}
	CommitMessageNegativeRegex *string `field:"optional" json:"commitMessageNegativeRegex" yaml:"commitMessageNegativeRegex"`
	// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#commit_message_regex Group#commit_message_regex}
	CommitMessageRegex *string `field:"optional" json:"commitMessageRegex" yaml:"commitMessageRegex"`
	// Deny deleting a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#deny_delete_tag Group#deny_delete_tag}
	DenyDeleteTag interface{} `field:"optional" json:"denyDeleteTag" yaml:"denyDeleteTag"`
	// Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#file_name_regex Group#file_name_regex}
	FileNameRegex *string `field:"optional" json:"fileNameRegex" yaml:"fileNameRegex"`
	// Maximum file size (MB) allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#max_file_size Group#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Allows only GitLab users to author commits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#member_check Group#member_check}
	MemberCheck interface{} `field:"optional" json:"memberCheck" yaml:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#prevent_secrets Group#prevent_secrets}
	PreventSecrets interface{} `field:"optional" json:"preventSecrets" yaml:"preventSecrets"`
	// Only commits signed through GPG are allowed.  **Note** This attribute is only supported in GitLab versions >= 16.4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/group#reject_unsigned_commits Group#reject_unsigned_commits}
	RejectUnsignedCommits interface{} `field:"optional" json:"rejectUnsignedCommits" yaml:"rejectUnsignedCommits"`
}

