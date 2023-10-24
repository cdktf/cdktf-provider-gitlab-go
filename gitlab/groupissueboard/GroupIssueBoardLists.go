// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupissueboard


type GroupIssueBoardLists struct {
	// The ID of the label the list should be scoped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.5.0/docs/resources/group_issue_board#label_id GroupIssueBoard#label_id}
	LabelId *float64 `field:"optional" json:"labelId" yaml:"labelId"`
}
