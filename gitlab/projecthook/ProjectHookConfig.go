// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projecthook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectHookConfig struct {
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
	// The name or id of the project to add the hook to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#project ProjectHook#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#url ProjectHook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Invoke the hook for confidential issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#confidential_issues_events ProjectHook#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Invoke the hook for confidential note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#confidential_note_events ProjectHook#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Custom headers for the project webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#custom_headers ProjectHook#custom_headers}
	CustomHeaders interface{} `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// Custom webhook template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#custom_webhook_template ProjectHook#custom_webhook_template}
	CustomWebhookTemplate *string `field:"optional" json:"customWebhookTemplate" yaml:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#deployment_events ProjectHook#deployment_events}
	DeploymentEvents interface{} `field:"optional" json:"deploymentEvents" yaml:"deploymentEvents"`
	// Description of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#description ProjectHook#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable SSL verification when invoking the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#enable_ssl_verification ProjectHook#enable_ssl_verification}
	EnableSslVerification interface{} `field:"optional" json:"enableSslVerification" yaml:"enableSslVerification"`
	// Invoke the hook for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#issues_events ProjectHook#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// Invoke the hook for job events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#job_events ProjectHook#job_events}
	JobEvents interface{} `field:"optional" json:"jobEvents" yaml:"jobEvents"`
	// Invoke the hook for merge requests events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#merge_requests_events ProjectHook#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Name of the project webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#name ProjectHook#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Invoke the hook for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#note_events ProjectHook#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Invoke the hook for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#pipeline_events ProjectHook#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// Invoke the hook for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#push_events ProjectHook#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#push_events_branch_filter ProjectHook#push_events_branch_filter}
	PushEventsBranchFilter *string `field:"optional" json:"pushEventsBranchFilter" yaml:"pushEventsBranchFilter"`
	// Invoke the hook for release events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#releases_events ProjectHook#releases_events}
	ReleasesEvents interface{} `field:"optional" json:"releasesEvents" yaml:"releasesEvents"`
	// Invoke the hook for project access token expiry events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#resource_access_token_events ProjectHook#resource_access_token_events}
	ResourceAccessTokenEvents interface{} `field:"optional" json:"resourceAccessTokenEvents" yaml:"resourceAccessTokenEvents"`
	// Invoke the hook for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#tag_push_events ProjectHook#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#token ProjectHook#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Invoke the hook for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.10.0/docs/resources/project_hook#wiki_page_events ProjectHook#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}

