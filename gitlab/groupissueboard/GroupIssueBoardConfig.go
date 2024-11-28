// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupissueboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupIssueBoardConfig struct {
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
	// The ID or URL-encoded path of the group owned by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/group_issue_board#group GroupIssueBoard#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the board.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/group_issue_board#name GroupIssueBoard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of label names which the board should be scoped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/group_issue_board#labels GroupIssueBoard#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// lists block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/group_issue_board#lists GroupIssueBoard#lists}
	Lists interface{} `field:"optional" json:"lists" yaml:"lists"`
	// The milestone the board should be scoped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/group_issue_board#milestone_id GroupIssueBoard#milestone_id}
	MilestoneId *float64 `field:"optional" json:"milestoneId" yaml:"milestoneId"`
}

