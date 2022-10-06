package servicemicrosoftteams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceMicrosoftTeamsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#project ServiceMicrosoftTeams#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#webhook ServiceMicrosoftTeams#webhook}
	Webhook *string `field:"required" json:"webhook" yaml:"webhook"`
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default_and_protected”. The default value is “default”.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#branches_to_be_notified ServiceMicrosoftTeams#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Enable notifications for confidential issue events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#confidential_issues_events ServiceMicrosoftTeams#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#confidential_note_events ServiceMicrosoftTeams#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#id ServiceMicrosoftTeams#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable notifications for issue events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#issues_events ServiceMicrosoftTeams#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#merge_requests_events ServiceMicrosoftTeams#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#note_events ServiceMicrosoftTeams#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Send notifications for broken pipelines.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#notify_only_broken_pipelines ServiceMicrosoftTeams#notify_only_broken_pipelines}
	NotifyOnlyBrokenPipelines interface{} `field:"optional" json:"notifyOnlyBrokenPipelines" yaml:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#pipeline_events ServiceMicrosoftTeams#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#push_events ServiceMicrosoftTeams#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#tag_push_events ServiceMicrosoftTeams#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Enable notifications for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/service_microsoft_teams#wiki_page_events ServiceMicrosoftTeams#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}
