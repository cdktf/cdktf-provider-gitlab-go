// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationmicrosoftteams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationMicrosoftTeamsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#project ProjectIntegrationMicrosoftTeams#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#webhook ProjectIntegrationMicrosoftTeams#webhook}
	Webhook *string `field:"required" json:"webhook" yaml:"webhook"`
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default_and_protected”. The default value is “default”.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#branches_to_be_notified ProjectIntegrationMicrosoftTeams#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Enable notifications for confidential issue events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#confidential_issues_events ProjectIntegrationMicrosoftTeams#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#confidential_note_events ProjectIntegrationMicrosoftTeams#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#id ProjectIntegrationMicrosoftTeams#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable notifications for issue events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#issues_events ProjectIntegrationMicrosoftTeams#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#merge_requests_events ProjectIntegrationMicrosoftTeams#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#note_events ProjectIntegrationMicrosoftTeams#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#notify_only_broken_pipelines ProjectIntegrationMicrosoftTeams#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#pipeline_events ProjectIntegrationMicrosoftTeams#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#push_events ProjectIntegrationMicrosoftTeams#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#tag_push_events ProjectIntegrationMicrosoftTeams#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/project_integration_microsoft_teams#wiki_page_events ProjectIntegrationMicrosoftTeams#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}

