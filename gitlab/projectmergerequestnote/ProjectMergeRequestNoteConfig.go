// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectmergerequestnote

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectMergeRequestNoteConfig struct {
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
	// The body of the merge request note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#body ProjectMergeRequestNote#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// The IID of the merge request to add the note to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#merge_request_iid ProjectMergeRequestNote#merge_request_iid}
	MergeRequestIid *float64 `field:"required" json:"mergeRequestIid" yaml:"mergeRequestIid"`
	// The ID or path of the project to add the note to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#project ProjectMergeRequestNote#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The creation date of the merge request note.
	//
	// Using this field requires the token used with the provider to either be an Admin, or hava a Project or Group Owner role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#created_at ProjectMergeRequestNote#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Indicates if the merge request note is internal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#internal ProjectMergeRequestNote#internal}
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// The diff head SHA of the merge request when the note was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_merge_request_note#merge_request_diff_head_sha ProjectMergeRequestNote#merge_request_diff_head_sha}
	MergeRequestDiffHeadSha *string `field:"optional" json:"mergeRequestDiffHeadSha" yaml:"mergeRequestDiffHeadSha"`
}

