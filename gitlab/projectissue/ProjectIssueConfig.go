// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectissue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIssueConfig struct {
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
	// The name or ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#project ProjectIssue#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The title of the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#title ProjectIssue#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The IDs of the users to assign the issue to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#assignee_ids ProjectIssue#assignee_ids}
	AssigneeIds *[]*float64 `field:"optional" json:"assigneeIds" yaml:"assigneeIds"`
	// Set an issue to be confidential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#confidential ProjectIssue#confidential}
	Confidential interface{} `field:"optional" json:"confidential" yaml:"confidential"`
	// When the issue was created.
	//
	// Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#created_at ProjectIssue#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Whether the issue is deleted instead of closed during destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#delete_on_destroy ProjectIssue#delete_on_destroy}
	DeleteOnDestroy interface{} `field:"optional" json:"deleteOnDestroy" yaml:"deleteOnDestroy"`
	// The description of an issue. Limited to 1,048,576 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#description ProjectIssue#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the issue is locked for discussions or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#discussion_locked ProjectIssue#discussion_locked}
	DiscussionLocked interface{} `field:"optional" json:"discussionLocked" yaml:"discussionLocked"`
	// The ID of a discussion to resolve.
	//
	// This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge_request_to_resolve_discussions_of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#discussion_to_resolve ProjectIssue#discussion_to_resolve}
	DiscussionToResolve *string `field:"optional" json:"discussionToResolve" yaml:"discussionToResolve"`
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#due_date ProjectIssue#due_date}
	DueDate *string `field:"optional" json:"dueDate" yaml:"dueDate"`
	// The ID of the epic issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#epic_issue_id ProjectIssue#epic_issue_id}
	EpicIssueId *float64 `field:"optional" json:"epicIssueId" yaml:"epicIssueId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#id ProjectIssue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The internal ID of the project's issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#iid ProjectIssue#iid}
	Iid *float64 `field:"optional" json:"iid" yaml:"iid"`
	// The type of issue. Valid values are: `issue`, `incident`, `test_case`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#issue_type ProjectIssue#issue_type}
	IssueType *string `field:"optional" json:"issueType" yaml:"issueType"`
	// The labels of an issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#labels ProjectIssue#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The IID of a merge request in which to resolve all issues.
	//
	// This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#merge_request_to_resolve_discussions_of ProjectIssue#merge_request_to_resolve_discussions_of}
	MergeRequestToResolveDiscussionsOf *float64 `field:"optional" json:"mergeRequestToResolveDiscussionsOf" yaml:"mergeRequestToResolveDiscussionsOf"`
	// The global ID of a milestone to assign issue.
	//
	// To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#milestone_id ProjectIssue#milestone_id}
	MilestoneId *float64 `field:"optional" json:"milestoneId" yaml:"milestoneId"`
	// The state of the issue. Valid values are: `opened`, `closed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#state ProjectIssue#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#updated_at ProjectIssue#updated_at}
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// The weight of the issue. Valid values are greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_issue#weight ProjectIssue#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

