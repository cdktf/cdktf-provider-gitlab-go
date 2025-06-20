// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationtelegram

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationTelegramConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#confidential_issues_events ProjectIntegrationTelegram#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"required" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#confidential_note_events ProjectIntegrationTelegram#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"required" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Enable notifications for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#issues_events ProjectIntegrationTelegram#issues_events}
	IssuesEvents interface{} `field:"required" json:"issuesEvents" yaml:"issuesEvents"`
	// Enable notifications for merge requests events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#merge_requests_events ProjectIntegrationTelegram#merge_requests_events}
	MergeRequestsEvents interface{} `field:"required" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#note_events ProjectIntegrationTelegram#note_events}
	NoteEvents interface{} `field:"required" json:"noteEvents" yaml:"noteEvents"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#pipeline_events ProjectIntegrationTelegram#pipeline_events}
	PipelineEvents interface{} `field:"required" json:"pipelineEvents" yaml:"pipelineEvents"`
	// The ID or full path of the project to integrate with Telegram.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#project ProjectIntegrationTelegram#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#push_events ProjectIntegrationTelegram#push_events}
	PushEvents interface{} `field:"required" json:"pushEvents" yaml:"pushEvents"`
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#room ProjectIntegrationTelegram#room}
	Room *string `field:"required" json:"room" yaml:"room"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#tag_push_events ProjectIntegrationTelegram#tag_push_events}
	TagPushEvents interface{} `field:"required" json:"tagPushEvents" yaml:"tagPushEvents"`
	// The Telegram bot token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#token ProjectIntegrationTelegram#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#wiki_page_events ProjectIntegrationTelegram#wiki_page_events}
	WikiPageEvents interface{} `field:"required" json:"wikiPageEvents" yaml:"wikiPageEvents"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#branches_to_be_notified ProjectIntegrationTelegram#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/project_integration_telegram#notify_only_broken_pipelines ProjectIntegrationTelegram#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
}

