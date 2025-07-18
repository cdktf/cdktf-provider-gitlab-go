// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globallevelnotifications

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalLevelNotificationsConfig struct {
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
	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#close_issue GlobalLevelNotifications#close_issue}
	CloseIssue interface{} `field:"optional" json:"closeIssue" yaml:"closeIssue"`
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#close_merge_request GlobalLevelNotifications#close_merge_request}
	CloseMergeRequest interface{} `field:"optional" json:"closeMergeRequest" yaml:"closeMergeRequest"`
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#failed_pipeline GlobalLevelNotifications#failed_pipeline}
	FailedPipeline interface{} `field:"optional" json:"failedPipeline" yaml:"failedPipeline"`
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#fixed_pipeline GlobalLevelNotifications#fixed_pipeline}
	FixedPipeline interface{} `field:"optional" json:"fixedPipeline" yaml:"fixedPipeline"`
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#issue_due GlobalLevelNotifications#issue_due}
	IssueDue interface{} `field:"optional" json:"issueDue" yaml:"issueDue"`
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#level GlobalLevelNotifications#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#merge_merge_request GlobalLevelNotifications#merge_merge_request}
	MergeMergeRequest interface{} `field:"optional" json:"mergeMergeRequest" yaml:"mergeMergeRequest"`
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#merge_when_pipeline_succeeds GlobalLevelNotifications#merge_when_pipeline_succeeds}
	MergeWhenPipelineSucceeds interface{} `field:"optional" json:"mergeWhenPipelineSucceeds" yaml:"mergeWhenPipelineSucceeds"`
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#moved_project GlobalLevelNotifications#moved_project}
	MovedProject interface{} `field:"optional" json:"movedProject" yaml:"movedProject"`
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#new_issue GlobalLevelNotifications#new_issue}
	NewIssue interface{} `field:"optional" json:"newIssue" yaml:"newIssue"`
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#new_merge_request GlobalLevelNotifications#new_merge_request}
	NewMergeRequest interface{} `field:"optional" json:"newMergeRequest" yaml:"newMergeRequest"`
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#new_note GlobalLevelNotifications#new_note}
	NewNote interface{} `field:"optional" json:"newNote" yaml:"newNote"`
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#push_to_merge_request GlobalLevelNotifications#push_to_merge_request}
	PushToMergeRequest interface{} `field:"optional" json:"pushToMergeRequest" yaml:"pushToMergeRequest"`
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#reassign_issue GlobalLevelNotifications#reassign_issue}
	ReassignIssue interface{} `field:"optional" json:"reassignIssue" yaml:"reassignIssue"`
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#reassign_merge_request GlobalLevelNotifications#reassign_merge_request}
	ReassignMergeRequest interface{} `field:"optional" json:"reassignMergeRequest" yaml:"reassignMergeRequest"`
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#reopen_issue GlobalLevelNotifications#reopen_issue}
	ReopenIssue interface{} `field:"optional" json:"reopenIssue" yaml:"reopenIssue"`
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#reopen_merge_request GlobalLevelNotifications#reopen_merge_request}
	ReopenMergeRequest interface{} `field:"optional" json:"reopenMergeRequest" yaml:"reopenMergeRequest"`
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/global_level_notifications#success_pipeline GlobalLevelNotifications#success_pipeline}
	SuccessPipeline interface{} `field:"optional" json:"successPipeline" yaml:"successPipeline"`
}

