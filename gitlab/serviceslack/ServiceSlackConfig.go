// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceslack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceSlackConfig struct {
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
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#project ServiceSlack#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#webhook ServiceSlack#webhook}
	Webhook *string `field:"required" json:"webhook" yaml:"webhook"`
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default_and_protected".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#branches_to_be_notified ServiceSlack#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#confidential_issue_channel ServiceSlack#confidential_issue_channel}
	ConfidentialIssueChannel *string `field:"optional" json:"confidentialIssueChannel" yaml:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#confidential_issues_events ServiceSlack#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#confidential_note_channel ServiceSlack#confidential_note_channel}
	ConfidentialNoteChannel *string `field:"optional" json:"confidentialNoteChannel" yaml:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#confidential_note_events ServiceSlack#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#id ServiceSlack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the channel to receive issue events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#issue_channel ServiceSlack#issue_channel}
	IssueChannel *string `field:"optional" json:"issueChannel" yaml:"issueChannel"`
	// Enable notifications for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#issues_events ServiceSlack#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// The name of the channel to receive merge request events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#merge_request_channel ServiceSlack#merge_request_channel}
	MergeRequestChannel *string `field:"optional" json:"mergeRequestChannel" yaml:"mergeRequestChannel"`
	// Enable notifications for merge requests events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#merge_requests_events ServiceSlack#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// The name of the channel to receive note events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#note_channel ServiceSlack#note_channel}
	NoteChannel *string `field:"optional" json:"noteChannel" yaml:"noteChannel"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#note_events ServiceSlack#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#notify_only_broken_pipelines ServiceSlack#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
	// This parameter has been replaced with `branches_to_be_notified`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#notify_only_default_branch ServiceSlack#notify_only_default_branch}
	NotifyOnlyDefaultBranch interface{} `field:"optional" json:"notifyOnlyDefaultBranch" yaml:"notifyOnlyDefaultBranch"`
	// The name of the channel to receive pipeline events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#pipeline_channel ServiceSlack#pipeline_channel}
	PipelineChannel *string `field:"optional" json:"pipelineChannel" yaml:"pipelineChannel"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#pipeline_events ServiceSlack#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// The name of the channel to receive push events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#push_channel ServiceSlack#push_channel}
	PushChannel *string `field:"optional" json:"pushChannel" yaml:"pushChannel"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#push_events ServiceSlack#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// The name of the channel to receive tag push events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#tag_push_channel ServiceSlack#tag_push_channel}
	TagPushChannel *string `field:"optional" json:"tagPushChannel" yaml:"tagPushChannel"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#tag_push_events ServiceSlack#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Username to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#username ServiceSlack#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// The name of the channel to receive wiki page events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#wiki_page_channel ServiceSlack#wiki_page_channel}
	WikiPageChannel *string `field:"optional" json:"wikiPageChannel" yaml:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs/resources/service_slack#wiki_page_events ServiceSlack#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}

