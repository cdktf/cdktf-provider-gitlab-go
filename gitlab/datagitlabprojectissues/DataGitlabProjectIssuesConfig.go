// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectissues

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectIssuesConfig struct {
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
	// The name or id of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#project DataGitlabProjectIssues#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Return issues assigned to the given user id.
	//
	// Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#assignee_id DataGitlabProjectIssues#assignee_id}
	AssigneeId *float64 `field:"optional" json:"assigneeId" yaml:"assigneeId"`
	// Return issues assigned to the given username.
	//
	// Similar to assignee_id and mutually exclusive with assignee_id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#assignee_username DataGitlabProjectIssues#assignee_username}
	AssigneeUsername *string `field:"optional" json:"assigneeUsername" yaml:"assigneeUsername"`
	// Return issues created by the given user id. Combine with scope=all or scope=assigned_to_me.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#author_id DataGitlabProjectIssues#author_id}
	AuthorId *float64 `field:"optional" json:"authorId" yaml:"authorId"`
	// Filter confidential or public issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#confidential DataGitlabProjectIssues#confidential}
	Confidential interface{} `field:"optional" json:"confidential" yaml:"confidential"`
	// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#created_after DataGitlabProjectIssues#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#created_before DataGitlabProjectIssues#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month.
	//
	// Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next_month_and_previous_two_weeks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#due_date DataGitlabProjectIssues#due_date}
	DueDate *string `field:"optional" json:"dueDate" yaml:"dueDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#id DataGitlabProjectIssues#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Return only the issues having the given iid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#iids DataGitlabProjectIssues#iids}
	Iids *[]*float64 `field:"optional" json:"iids" yaml:"iids"`
	// Filter to a given type of issue. Valid values are [issue incident test_case].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#issue_type DataGitlabProjectIssues#issue_type}
	IssueType *string `field:"optional" json:"issueType" yaml:"issueType"`
	// Return issues with labels.
	//
	// Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#labels DataGitlabProjectIssues#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#milestone DataGitlabProjectIssues#milestone}
	Milestone *string `field:"optional" json:"milestone" yaml:"milestone"`
	// Return issues reacted by the authenticated user by the given emoji.
	//
	// None returns issues not given a reaction. Any returns issues given at least one reaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#my_reaction_emoji DataGitlabProjectIssues#my_reaction_emoji}
	MyReactionEmoji *string `field:"optional" json:"myReactionEmoji" yaml:"myReactionEmoji"`
	// Return issues that do not match the assignee id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#not_assignee_id DataGitlabProjectIssues#not_assignee_id}
	NotAssigneeId *float64 `field:"optional" json:"notAssigneeId" yaml:"notAssigneeId"`
	// Return issues that do not match the author id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#not_author_id DataGitlabProjectIssues#not_author_id}
	NotAuthorId *float64 `field:"optional" json:"notAuthorId" yaml:"notAuthorId"`
	// Return issues that do not match the labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#not_labels DataGitlabProjectIssues#not_labels}
	NotLabels *[]*string `field:"optional" json:"notLabels" yaml:"notLabels"`
	// Return issues that do not match the milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#not_milestone DataGitlabProjectIssues#not_milestone}
	NotMilestone *string `field:"optional" json:"notMilestone" yaml:"notMilestone"`
	// Return issues not reacted by the authenticated user by the given emoji.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#not_my_reaction_emoji DataGitlabProjectIssues#not_my_reaction_emoji}
	NotMyReactionEmoji *string `field:"optional" json:"notMyReactionEmoji" yaml:"notMyReactionEmoji"`
	// Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#order_by DataGitlabProjectIssues#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#scope DataGitlabProjectIssues#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Search project issues against their title and description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#search DataGitlabProjectIssues#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Return issues sorted in asc or desc order. Default is desc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#sort DataGitlabProjectIssues#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#updated_after DataGitlabProjectIssues#updated_after}
	UpdatedAfter *string `field:"optional" json:"updatedAfter" yaml:"updatedAfter"`
	// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#updated_before DataGitlabProjectIssues#updated_before}
	UpdatedBefore *string `field:"optional" json:"updatedBefore" yaml:"updatedBefore"`
	// Return issues with the specified weight.
	//
	// None returns issues with no weight assigned. Any returns issues with a weight assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#weight DataGitlabProjectIssues#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
	// If true, the response returns more details for each label in labels field: :name, :color, :description, :description_html, :text_color.
	//
	// Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/data-sources/project_issues#with_labels_details DataGitlabProjectIssues#with_labels_details}
	WithLabelsDetails interface{} `field:"optional" json:"withLabelsDetails" yaml:"withLabelsDetails"`
}

