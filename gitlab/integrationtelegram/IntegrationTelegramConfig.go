// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationtelegram

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationTelegramConfig struct {
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
	// Enable notifications for confidential issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#confidential_issues_events IntegrationTelegram#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"required" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#confidential_note_events IntegrationTelegram#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"required" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Enable notifications for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#issues_events IntegrationTelegram#issues_events}
	IssuesEvents interface{} `field:"required" json:"issuesEvents" yaml:"issuesEvents"`
	// Enable notifications for merge requests events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#merge_requests_events IntegrationTelegram#merge_requests_events}
	MergeRequestsEvents interface{} `field:"required" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#note_events IntegrationTelegram#note_events}
	NoteEvents interface{} `field:"required" json:"noteEvents" yaml:"noteEvents"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#pipeline_events IntegrationTelegram#pipeline_events}
	PipelineEvents interface{} `field:"required" json:"pipelineEvents" yaml:"pipelineEvents"`
	// The ID or full path of the project to integrate with Telegram.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#project IntegrationTelegram#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#push_events IntegrationTelegram#push_events}
	PushEvents interface{} `field:"required" json:"pushEvents" yaml:"pushEvents"`
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#room IntegrationTelegram#room}
	Room *string `field:"required" json:"room" yaml:"room"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#tag_push_events IntegrationTelegram#tag_push_events}
	TagPushEvents interface{} `field:"required" json:"tagPushEvents" yaml:"tagPushEvents"`
	// The Telegram bot token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#token IntegrationTelegram#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#wiki_page_events IntegrationTelegram#wiki_page_events}
	WikiPageEvents interface{} `field:"required" json:"wikiPageEvents" yaml:"wikiPageEvents"`
	// Branches to send notifications for (introduced in GitLab 16.5). Update of this attribute was not supported before Gitlab 16.11 due to API bug. Valid options are `all`, `default`, `protected`, `default_and_protected`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#branches_to_be_notified IntegrationTelegram#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_telegram#notify_only_broken_pipelines IntegrationTelegram#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
}

