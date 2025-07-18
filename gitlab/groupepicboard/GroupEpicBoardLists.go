// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupepicboard


type GroupEpicBoardLists struct {
	// The ID of the label the list should be scoped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group_epic_board#label_id GroupEpicBoard#label_id}
	LabelId *float64 `field:"optional" json:"labelId" yaml:"labelId"`
}

