// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectmergerequests

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectMergeRequestsConfig struct {
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
	// The ID or path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#project DataGitlabProjectMergeRequests#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Return merge requests created by the given user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#author_id DataGitlabProjectMergeRequests#author_id}
	AuthorId *float64 `field:"optional" json:"authorId" yaml:"authorId"`
	// Return merge requests created by the given username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#author_username DataGitlabProjectMergeRequests#author_username}
	AuthorUsername *string `field:"optional" json:"authorUsername" yaml:"authorUsername"`
	// Return merge requests created after the given time. Expected in RFC3339 format (2006-01-02T15:04:05Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#created_after DataGitlabProjectMergeRequests#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Return merge requests created before the given time. Expected in RFC3339 format (2006-01-02T15:04:05Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#created_before DataGitlabProjectMergeRequests#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// The unique internal IDs of the merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#iids DataGitlabProjectMergeRequests#iids}
	Iids *[]*float64 `field:"optional" json:"iids" yaml:"iids"`
	// Return only merge requests for a specific milestone.
	//
	// `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#milestone DataGitlabProjectMergeRequests#milestone}
	Milestone *string `field:"optional" json:"milestone" yaml:"milestone"`
	// Return merge requests reacted to by the authenticated user with the given emoji.
	//
	// `None` returns issues not given a reaction. `Any` returns issues given at least one reaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#my_reaction_emoji DataGitlabProjectMergeRequests#my_reaction_emoji}
	MyReactionEmoji *string `field:"optional" json:"myReactionEmoji" yaml:"myReactionEmoji"`
	// Return requests ordered by `created_at`, `title` or `updated_at`. Default is `created_at`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#order_by DataGitlabProjectMergeRequests#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Return merge requests reviewed by the given username.
	//
	// `None` returns merge requests with no reviews. `Any` returns merge requests with any reviewer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#reviewer_username DataGitlabProjectMergeRequests#reviewer_username}
	ReviewerUsername *string `field:"optional" json:"reviewerUsername" yaml:"reviewerUsername"`
	// Return merge requests for the given scope: `created_by_me`, `assigned_to_me`, or `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#scope DataGitlabProjectMergeRequests#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Search merge requests against their `title` or `description`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#search DataGitlabProjectMergeRequests#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Return requests sorted in `asc` or `desc` order. Default is `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#sort DataGitlabProjectMergeRequests#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Return merge requests with the given source branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#source_branch DataGitlabProjectMergeRequests#source_branch}
	SourceBranch *string `field:"optional" json:"sourceBranch" yaml:"sourceBranch"`
	// Return all merge requests (all) or just those that are opened, closed, locked, or merged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#state DataGitlabProjectMergeRequests#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Return merge requests with the given target branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#target_branch DataGitlabProjectMergeRequests#target_branch}
	TargetBranch *string `field:"optional" json:"targetBranch" yaml:"targetBranch"`
	// Return merge requests updated after the given time. Expected in RFC3339 format (2006-01-02T15:04:05Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#updated_after DataGitlabProjectMergeRequests#updated_after}
	UpdatedAfter *string `field:"optional" json:"updatedAfter" yaml:"updatedAfter"`
	// Return merge requests updated before the given time. Expected in RFC3339 format (2006-01-02T15:04:05Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#updated_before DataGitlabProjectMergeRequests#updated_before}
	UpdatedBefore *string `field:"optional" json:"updatedBefore" yaml:"updatedBefore"`
	// Filter merge requests against their wip status.
	//
	// `yes` to return only draft merge requests, `no` to return non-draft merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/data-sources/project_merge_requests#wip DataGitlabProjectMergeRequests#wip}
	Wip *string `field:"optional" json:"wip" yaml:"wip"`
}

