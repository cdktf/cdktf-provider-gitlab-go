// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationmattermost

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationMattermostConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#project ProjectIntegrationMattermost#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#webhook ProjectIntegrationMattermost#webhook}
	Webhook *string `field:"required" json:"webhook" yaml:"webhook"`
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default_and_protected".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#branches_to_be_notified ProjectIntegrationMattermost#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#confidential_issue_channel ProjectIntegrationMattermost#confidential_issue_channel}
	ConfidentialIssueChannel *string `field:"optional" json:"confidentialIssueChannel" yaml:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#confidential_issues_events ProjectIntegrationMattermost#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#confidential_note_channel ProjectIntegrationMattermost#confidential_note_channel}
	ConfidentialNoteChannel *string `field:"optional" json:"confidentialNoteChannel" yaml:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#confidential_note_events ProjectIntegrationMattermost#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#id ProjectIntegrationMattermost#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the channel to receive issue events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#issue_channel ProjectIntegrationMattermost#issue_channel}
	IssueChannel *string `field:"optional" json:"issueChannel" yaml:"issueChannel"`
	// Enable notifications for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#issues_events ProjectIntegrationMattermost#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// The name of the channel to receive merge request events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#merge_request_channel ProjectIntegrationMattermost#merge_request_channel}
	MergeRequestChannel *string `field:"optional" json:"mergeRequestChannel" yaml:"mergeRequestChannel"`
	// Enable notifications for merge requests events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#merge_requests_events ProjectIntegrationMattermost#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// The name of the channel to receive note events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#note_channel ProjectIntegrationMattermost#note_channel}
	NoteChannel *string `field:"optional" json:"noteChannel" yaml:"noteChannel"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#note_events ProjectIntegrationMattermost#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#notify_only_broken_pipelines ProjectIntegrationMattermost#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
	// The name of the channel to receive pipeline events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#pipeline_channel ProjectIntegrationMattermost#pipeline_channel}
	PipelineChannel *string `field:"optional" json:"pipelineChannel" yaml:"pipelineChannel"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#pipeline_events ProjectIntegrationMattermost#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// The name of the channel to receive push events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#push_channel ProjectIntegrationMattermost#push_channel}
	PushChannel *string `field:"optional" json:"pushChannel" yaml:"pushChannel"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#push_events ProjectIntegrationMattermost#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// The name of the channel to receive tag push events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#tag_push_channel ProjectIntegrationMattermost#tag_push_channel}
	TagPushChannel *string `field:"optional" json:"tagPushChannel" yaml:"tagPushChannel"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#tag_push_events ProjectIntegrationMattermost#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Username to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#username ProjectIntegrationMattermost#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// The name of the channel to receive wiki page events notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#wiki_page_channel ProjectIntegrationMattermost#wiki_page_channel}
	WikiPageChannel *string `field:"optional" json:"wikiPageChannel" yaml:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_integration_mattermost#wiki_page_events ProjectIntegrationMattermost#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}

