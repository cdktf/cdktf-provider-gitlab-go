package project


type ProjectPushRules struct {
	// All commit author emails must match this regex, e.g. `@my-company.com$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#author_email_regex Project#author_email_regex}
	AuthorEmailRegex *string `field:"optional" json:"authorEmailRegex" yaml:"authorEmailRegex"`
	// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#branch_name_regex Project#branch_name_regex}
	BranchNameRegex *string `field:"optional" json:"branchNameRegex" yaml:"branchNameRegex"`
	// Users can only push commits to this repository that were committed with one of their own verified emails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#commit_committer_check Project#commit_committer_check}
	CommitCommitterCheck interface{} `field:"optional" json:"commitCommitterCheck" yaml:"commitCommitterCheck"`
	// No commit message is allowed to match this regex, for example `ssh\:\/\/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#commit_message_negative_regex Project#commit_message_negative_regex}
	CommitMessageNegativeRegex *string `field:"optional" json:"commitMessageNegativeRegex" yaml:"commitMessageNegativeRegex"`
	// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#commit_message_regex Project#commit_message_regex}
	CommitMessageRegex *string `field:"optional" json:"commitMessageRegex" yaml:"commitMessageRegex"`
	// Deny deleting a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#deny_delete_tag Project#deny_delete_tag}
	DenyDeleteTag interface{} `field:"optional" json:"denyDeleteTag" yaml:"denyDeleteTag"`
	// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#file_name_regex Project#file_name_regex}
	FileNameRegex *string `field:"optional" json:"fileNameRegex" yaml:"fileNameRegex"`
	// Maximum file size (MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#max_file_size Project#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Restrict commits by author (email) to existing GitLab users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#member_check Project#member_check}
	MemberCheck interface{} `field:"optional" json:"memberCheck" yaml:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#prevent_secrets Project#prevent_secrets}
	PreventSecrets interface{} `field:"optional" json:"preventSecrets" yaml:"preventSecrets"`
	// Reject commit when it’s not signed through GPG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.0/docs/resources/project#reject_unsigned_commits Project#reject_unsigned_commits}
	RejectUnsignedCommits interface{} `field:"optional" json:"rejectUnsignedCommits" yaml:"rejectUnsignedCommits"`
}

