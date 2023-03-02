package branchprotection


type BranchProtectionAllowedToUnprotect struct {
	// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/branch_protection#group_id BranchProtection#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/branch_protection#user_id BranchProtection#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

