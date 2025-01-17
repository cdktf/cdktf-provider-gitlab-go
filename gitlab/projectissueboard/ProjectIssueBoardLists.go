// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectissueboard


type ProjectIssueBoardLists struct {
	// The ID of the assignee the list should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_issue_board#assignee_id ProjectIssueBoard#assignee_id}
	AssigneeId *float64 `field:"optional" json:"assigneeId" yaml:"assigneeId"`
	// The ID of the iteration the list should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_issue_board#iteration_id ProjectIssueBoard#iteration_id}
	IterationId *float64 `field:"optional" json:"iterationId" yaml:"iterationId"`
	// The ID of the label the list should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_issue_board#label_id ProjectIssueBoard#label_id}
	LabelId *float64 `field:"optional" json:"labelId" yaml:"labelId"`
	// The ID of the milestone the list should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/project_issue_board#milestone_id ProjectIssueBoard#milestone_id}
	MilestoneId *float64 `field:"optional" json:"milestoneId" yaml:"milestoneId"`
}

