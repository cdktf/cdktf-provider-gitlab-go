// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projecttargetbranchrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectTargetBranchRuleConfig struct {
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
	// The ID or URL-encoded path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/project_target_branch_rule#project ProjectTargetBranchRule#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// A pattern matching the branch name for which the merge request should have a default target branch configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/project_target_branch_rule#source_branch_pattern ProjectTargetBranchRule#source_branch_pattern}
	SourceBranchPattern *string `field:"required" json:"sourceBranchPattern" yaml:"sourceBranchPattern"`
	// The name of the branch to which the merge request should be addressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/project_target_branch_rule#target_branch_name ProjectTargetBranchRule#target_branch_name}
	TargetBranchName *string `field:"required" json:"targetBranchName" yaml:"targetBranchName"`
}

