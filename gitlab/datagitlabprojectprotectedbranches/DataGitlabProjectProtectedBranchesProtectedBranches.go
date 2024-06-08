// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedbranches


type DataGitlabProjectProtectedBranchesProtectedBranches struct {
	// merge_access_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_protected_branches#merge_access_levels DataGitlabProjectProtectedBranches#merge_access_levels}
	MergeAccessLevels interface{} `field:"optional" json:"mergeAccessLevels" yaml:"mergeAccessLevels"`
	// push_access_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/data-sources/project_protected_branches#push_access_levels DataGitlabProjectProtectedBranches#push_access_levels}
	PushAccessLevels interface{} `field:"optional" json:"pushAccessLevels" yaml:"pushAccessLevels"`
}

